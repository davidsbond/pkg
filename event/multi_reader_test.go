package event_test

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/closers"
	"pkg.dsb.dev/event"
)

func TestMultiReader_Read(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())

	aWriter, err := event.NewWriter(ctx, "mem://a")
	assert.NoError(t, err)
	defer closers.Close(aWriter)

	bWriter, err := event.NewWriter(ctx, "mem://b")
	assert.NoError(t, err)
	defer closers.Close(bWriter)

	cWriter, err := event.NewWriter(ctx, "mem://c")
	assert.NoError(t, err)
	defer closers.Close(cWriter)

	rd, err := event.NewMultiReader(ctx, []string{"mem://a", "mem://b", "mem://c"})
	assert.NoError(t, err)
	defer closers.Close(rd)

	assert.NoError(t, aWriter.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))
	assert.NoError(t, bWriter.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))
	assert.NoError(t, cWriter.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))

	i := 0
	mux := &sync.Mutex{}
	err = rd.Read(ctx, func(ctx context.Context, evt event.Event) error {
		mux.Lock()
		defer mux.Unlock()

		assert.EqualValues(t, []byte("hello world"), evt.Payload)
		i++
		if i == 3 {
			cancel()
		}

		return nil
	})
	assert.True(t, errors.Is(err, context.Canceled))
}
