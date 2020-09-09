package event_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/closers"
	"pkg.dsb.dev/event"
)

func TestReader_Read(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	wr, err := event.NewWriter(ctx, "mem://test-topic")
	assert.NoError(t, err)
	defer closers.Close(wr)

	rd, err := event.NewReader(ctx, "mem://test-topic")
	assert.NoError(t, err)
	defer closers.Close(rd)

	assert.NoError(t, wr.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))

	err = rd.Read(ctx, func(ctx context.Context, evt event.Event) error {
		assert.EqualValues(t, []byte("hello world"), evt.Payload)
		cancel()
		return nil
	})
	assert.True(t, errors.Is(err, context.Canceled))
}
