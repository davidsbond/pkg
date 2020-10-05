package badger

import (
	"time"

	"github.com/dgraph-io/badger/v2"
)

type (
	// The Option type is a function that modifies the badger configuration.
	Option func(opts *badger.Options)
)

// WithDir sets the location on disk badger will write data.
func WithDir(dir string) Option {
	return func(opts *badger.Options) {
		opts.Dir = dir
		opts.ValueDir = dir
	}
}

// WithEncryptionKey sets the key to use to encrypt data at rest on the
// filesystem.
func WithEncryptionKey(key []byte) Option {
	return func(opts *badger.Options) {
		opts.EncryptionKey = key
	}
}

// WithEncryptionKeyRotationDuration sets how often to change the generated
// encryption key.
func WithEncryptionKeyRotationDuration(dur time.Duration) Option {
	return func(opts *badger.Options) {
		opts.EncryptionKeyRotationDuration = dur
	}
}
