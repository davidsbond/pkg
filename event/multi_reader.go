package event

import (
	"context"

	"golang.org/x/sync/errgroup"

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
func (mr *MultiReader) Read(ctx context.Context, fn Handler) error {
	grp, ctx := errgroup.WithContext(ctx)
	for _, rd := range mr.readers {
		mr.work(ctx, grp, rd, fn)
	}

	return grp.Wait()
}

func (mr *MultiReader) work(ctx context.Context, grp *errgroup.Group, rd *Reader, fn Handler) {
	grp.Go(func() error {
		return rd.Read(ctx, fn)
	})
}

// Close all event stream connections.
func (mr *MultiReader) Close() error {
	for _, rd := range mr.readers {
		closers.Close(rd)
	}
	return nil
}
