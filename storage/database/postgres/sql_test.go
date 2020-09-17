package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/environment"
	"pkg.dsb.dev/storage/database/postgres"
	"pkg.dsb.dev/testutil"
)

func TestCreateDatabaseWithUser(t *testing.T) {
	db := testutil.WithPostgresInstance(t)
	ctx := environment.NewContext()

	user := "test"
	dbName := "test"
	pass := "test"

	err := postgres.CreateDatabaseWithUser(ctx, db, dbName, user, pass)
	assert.NoError(t, err)
}
