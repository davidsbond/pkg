package sqlite_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/testutil"
)

func TestOpen(t *testing.T) {
	db := testutil.WithSQLiteInstance(t, nil)

	assert.NotNil(t, db)
	assert.FileExists(t, "test.db")
}
