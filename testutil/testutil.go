// Package testutil contains utility methods for use in tests
package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/storage/database/postgres"
)

// WithPostgresInstance is a test helper function that creates a connection to a postgres
// database configured by the environment. Once connected, migrations are performed. When the
// test is finished, the database will be migrated down again.
func WithPostgresInstance(t *testing.T) *sql.DB {
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")

	if pgHost == "" {
		pgHost = "localhost"
	}

	if pgPort == "" {
		pgPort = "5432"
	}

	if pgUser == "" {
		pgUser = "postgres"
	}

	if pgPass == "" {
		pgPass = "postgres"
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", pgUser, pgPass, pgHost, pgPort)

	db, err := postgres.Open(url, nil)
	if err != nil {
		assert.FailNow(t, err.Error())
		return nil
	}

	dbName := xid.New().String()
	_, err = db.Exec("CREATE DATABASE " + dbName)
	assert.NoError(t, err)
	assert.NoError(t, db.Close())

	url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPass, pgHost, pgPort, dbName)
	db, err = postgres.Open(url, nil)
	if err != nil {
		assert.FailNow(t, err.Error())
		return nil
	}

	t.Cleanup(func() {
		assert.NoError(t, db.Close())
	})

	return db
}

type (
	// The Matcher type is a generic gomock.Matcher implementation for comparing two
	// values.
	Matcher struct {
		t        *testing.T
		expected interface{}
	}
)

// Matches returns a generic gomock.Matcher implementation for a value.
func Matches(t *testing.T, exp interface{}) gomock.Matcher {
	return &Matcher{
		t:        t,
		expected: exp,
	}
}

// Matches returns true if the given value matches the expected one.
func (m *Matcher) Matches(actual interface{}) bool {
	return assert.EqualValues(m.t, m.expected, actual)
}

func (m *Matcher) String() string {
	return fmt.Sprintf("is equal to %v", m.expected)
}
