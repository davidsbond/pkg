package event_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/closers"
	"pkg.dsb.dev/event"
)

func TestWriter_Write(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	wr, err := event.NewWriter(ctx, "mem://test-topic")
	defer closers.Close(wr)

	assert.NoError(t, err)
	assert.NoError(t, wr.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))
}
