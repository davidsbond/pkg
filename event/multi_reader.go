package event

import (
	"context"

	"pkg.dsb.dev/closers"
)

type (
	// The MultiReader type is used to handle inbound events from multiple topics across different
	// event stream providers.
	MultiReader struct {
		readers []*Reader
	}
)

// NewMultiReader creates a new instance of the MultiReader type that will read events from the configured
// event stream providers identified using the given URLs.
func NewMultiReader(ctx context.Context, urls []string) (*MultiReader, error) {
	readers := make([]*Reader, 0)
	for _, url := range urls {
		reader, err := NewReader(ctx, url)
		if err != nil {
			for _, rd := range readers {
				closers.Close(rd)
			}
			return nil, err
		}
		readers = append(readers, reader)
	}

	return &MultiReader{readers: readers}, nil
}

// Read events from the stream, invoking fn for each inbound event. This method will block until fn returns
// an error or the provided context is cancelled.
func (mr *MultiReader) Read(ctx context.Context, fn func(ctx context.Context, evt Event) error) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	errs := make(chan error, len(mr.readers))

	for _, rd := range mr.readers {
		go func(r *Reader) {
			errs <- r.Read(ctx, fn)
		}(rd)
	}

	return <-errs
}

// Close all event stream connections.
func (mr *MultiReader) Close() error {
	for _, rd := range mr.readers {
		closers.Close(rd)
	}
	return nil
}
