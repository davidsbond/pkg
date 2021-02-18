package event_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"pkg.dsb.dev/closers"
	"pkg.dsb.dev/event"
)

func TestMultiReader_Read(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	payload := timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	expected := event.New(payload)

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

	assert.NoError(t, aWriter.Write(ctx, expected))
	assert.NoError(t, bWriter.Write(ctx, expected))
	assert.NoError(t, cWriter.Write(ctx, expected))

	i := 0
	mux := &sync.Mutex{}
	err = rd.Read(ctx, func(ctx context.Context, evt event.Event) error {
		mux.Lock()
		defer mux.Unlock()

		assert.EqualValues(t, expected.ID, evt.ID)
		i++
		if i == 3 {
			cancel()
		}

		return nil
	})
	assert.True(t, errors.Is(err, context.Canceled))
}
