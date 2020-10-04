// Package badger provides an instrumented wrapper around a badger database.
package badger

import (
	"context"

	"github.com/dgraph-io/badger/v2"
	"github.com/opentracing/opentracing-go"

	"pkg.dsb.dev/health"
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

// Open a badger database using the provided options.
func Open(options badger.Options) (*DB, error) {
	inner, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	db := &DB{inner: inner}
	health.AddCheck(options.Dir, db.Ping)
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
