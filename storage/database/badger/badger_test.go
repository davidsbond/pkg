package badger_test

import (
	"context"
	"os"
	"testing"

	bdg "github.com/dgraph-io/badger/v2"
	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/storage/database/badger"
)

func TestOpen(t *testing.T) {
	db, err := badger.Open(bdg.DefaultOptions("test"))
	assert.NoError(t, err)
	assert.NotNil(t, db)
	t.Cleanup(func() {
		assert.NoError(t, db.Close())
		assert.NoError(t, os.RemoveAll("test"))
	})

	ctx := context.Background()
	assert.NoError(t, db.Update(ctx, func(ctx context.Context, txn *badger.Txn) error {
		return txn.Set(ctx, []byte("hello"), []byte("world"))
	}))

	assert.NoError(t, db.View(ctx, func(ctx context.Context, txn *badger.Txn) error {
		item, err := txn.Get(ctx, []byte("hello"))
		if err != nil {
			return err
		}

		return item.Value(ctx, func(ctx context.Context, value []byte) error {
			assert.EqualValues(t, []byte("world"), value)
			return nil
		})
	}))
}
