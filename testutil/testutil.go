// Package testutil contains utility methods for use in tests
package testutil

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/storage/database"
	"pkg.dsb.dev/storage/database/postgres"
	"pkg.dsb.dev/storage/database/sqlite"
	"pkg.dsb.dev/storage/ftp"
)

// WithSQLiteInstance is a test helper function that creates an sqlite database.
// Once connected, migrations are performed. When the test is finished, the database
// will be migrated down again.
func WithSQLiteInstance(t *testing.T, migrations *database.MigrationSource) *sql.DB {
	t.Helper()

	const url = "test.db"

	db, err := sqlite.Open(url, migrations)
	if err != nil {
		assert.FailNow(t, err.Error())
		return nil
	}

	t.Cleanup(func() {
		if migrations != nil {
			assert.NoError(t, database.MigrateDown(migrations, db))
		}

		assert.NoError(t, db.Close())
		assert.NoError(t, os.Remove("test.db"))
	})

	return db
}

// WithPostgresInstance is a test helper function that creates a connection to a postgres
// database configured by the environment. Once connected, migrations are performed. When the
// test is finished, the database will be migrated down again.
func WithPostgresInstance(t *testing.T, migrations *database.MigrationSource) *sql.DB {
	t.Helper()

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

	t.Cleanup(func() {
		if migrations != nil {
			assert.NoError(t, database.MigrateDown(migrations, db))
		}

		assert.NoError(t, db.Close())
	})

	return db
}

// WithFTPServer is a test helper function that creates a connection to an FTP
// server configured by the environment.
func WithFTPServer(t *testing.T) *ftp.Conn {
	t.Helper()

	ftpHost := os.Getenv("FTP_HOST")
	ftpPort := os.Getenv("FTP_PORT")
	ftpUser := os.Getenv("FTP_USER")
	ftpPass := os.Getenv("FTP_PASSWORD")

	if ftpHost == "" {
		ftpHost = "localhost"
	}

	if ftpPort == "" {
		ftpPort = "21"
	}

	if ftpUser == "" {
		ftpUser = "ftp"
	}

	if ftpPass == "" {
		ftpPass = "ftp"
	}

	url := fmt.Sprintf("%s:%s", ftpHost, ftpPort)
	conn, err := ftp.Open(context.Background(), url, ftp.WithCredentials(ftpUser, ftpPass))
	if err != nil {
		assert.FailNow(t, err.Error())
		return nil
	}

	t.Cleanup(func() {
		assert.NoError(t, conn.Close())
	})

	return conn
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
	t.Helper()
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
