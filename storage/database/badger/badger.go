// Package badger provides an instrumented wrapper around a badger database.
package badger

import (
	"context"
	"io"

	"github.com/dgraph-io/badger/v2"
	"github.com/opentracing/opentracing-go"

	"pkg.dsb.dev/health"
	"pkg.dsb.dev/logging"
	"pkg.dsb.dev/tracing"
)

type (
	// The DB type represents the badger DB connection and is the main entrypoint
	// for querying and manipulating data.
	DB struct {
		inner *badger.DB
	}

	// The Txn type represents a database transaction, and is used to query and modify
	// data.
	Txn struct {
		inner *badger.Txn
	}

	// The Item type represents a key/value item stored in the database.
	Item struct {
		inner *badger.Item
	}
)

// Open a badger database using the provided options. Uses badger.DefaultOptions
// storing data in a "badger" directory.
func Open(opts ...Option) (*DB, error) {
	// Default directory is named "badger" and logging is
	// disabled.
	c := badger.DefaultOptions("badger")
	c.Logger = logging.Logger()

	for _, opt := range opts {
		opt(&c)
	}

	inner, err := badger.Open(c)
	if err != nil {
		return nil, err
	}

	db := &DB{inner: inner}
	health.AddCheck(c.Dir, db.Ping)
	return db, nil
}

// Ping syncs the badger database to determine everything is working as
// expected.
func (db *DB) Ping() error {
	return db.inner.Sync()
}

// Close the connection to the database.
func (db *DB) Close() error {
	return db.inner.Close()
}

// Backup writes a full backup to the provided io.Writer implementation.
func (db *DB) Backup(ctx context.Context, wr io.Writer) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "badger-backup")
	defer span.Finish()

	n, err := db.inner.Backup(wr, 0)
	if err != nil {
		return tracing.WithError(span, err)
	}

	span.SetTag("backup.timestamp", n)
	return nil
}

// Restore the database to the dump provided in the io.Reader implementation.
func (db *DB) Restore(ctx context.Context, rd io.Reader) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "badger-restore")
	defer span.Finish()

	const maxPending = 10
	return tracing.WithError(span, db.inner.Load(rd, maxPending))
}

// View executes a function creating and managing a read-only transaction for the user. Error
// returned by the function is relayed by the View method.
// If View is used with managed transactions, it would assume a read timestamp of MaxUint64.
func (db *DB) View(ctx context.Context, fn func(ctx context.Context, txn *Txn) error) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "badger-view")
	defer span.Finish()

	return db.inner.View(func(txn *badger.Txn) error {
		span.SetTag("txn.read_ts", txn.ReadTs())
		return tracing.WithError(span, fn(ctx, &Txn{inner: txn}))
	})
}

// Update executes a function, creating and managing a read-write transaction
// for the user. Error returned by the function is relayed by the Update method.
// Update cannot be used with managed transactions.
func (db *DB) Update(ctx context.Context, fn func(ctx context.Context, txn *Txn) error) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "badger-update")
	defer span.Finish()

	return db.inner.Update(func(txn *badger.Txn) error {
		span.SetTag("txn.read_ts", txn.ReadTs())

		return tracing.WithError(span, fn(ctx, &Txn{inner: txn}))
	})
}

// Get looks for key and returns corresponding Item.
// If key is not found, ErrKeyNotFound is returned.
func (txn *Txn) Get(ctx context.Context, key []byte) (*Item, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "badger-txn-get")
	defer span.Finish()
	span.SetTag("item.key", key)

	item, err := txn.inner.Get(key)
	if err != nil {
		return nil, tracing.WithError(span, err)
	}

	span.SetTag("item.version", item.Version())
	span.SetTag("item.key_size", item.KeySize())
	span.SetTag("item.value_size", item.ValueSize())
	return &Item{inner: item}, nil
}

// Set adds a key-value pair to the database.
// It will return ErrReadOnlyTxn if update flag was set to false when creating the transaction.
//
// The current transaction keeps a reference to the key and val byte slice
// arguments. Users must not modify key and val until the end of the transaction.
func (txn *Txn) Set(ctx context.Context, key, value []byte) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "badger-txn-set")
	defer span.Finish()
	span.SetTag("item.key", key)

	err := txn.inner.Set(key, value)
	return tracing.WithError(span, err)
}

// Delete deletes a key.
//
// This is done by adding a delete marker for the key at commit timestamp.  Any
// reads happening before this timestamp would be unaffected. Any reads after
// this commit would see the deletion.
//
// The current transaction keeps a reference to the key byte slice argument.
// Users must not modify the key until the end of the transaction.
func (txn *Txn) Delete(ctx context.Context, key []byte) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "badger-txn-delete")
	defer span.Finish()
	span.SetTag("item.key", key)

	err := txn.inner.Delete(key)
	return tracing.WithError(span, err)
}

// Value retrieves the value of the item from the value log.
//
// This method must be called within a transaction. Calling it outside a
// transaction is considered undefined behavior. If an iterator is being used,
// then Item.Value() is defined in the current iteration only, because items are
// reused.
//
// If you need to use a value outside a transaction, please use Item.ValueCopy
// instead, or copy it yourself. Value might change once discard or commit is called.
// Use ValueCopy if you want to do a Set after Get.
func (item *Item) Value(ctx context.Context, fn func(ctx context.Context, value []byte) error) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "badger-item-value")
	defer span.Finish()

	return tracing.WithError(span, item.inner.Value(func(val []byte) error {
		return fn(ctx, val)
	}))
}
