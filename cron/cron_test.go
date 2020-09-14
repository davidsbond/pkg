package cron_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/cron"
)

func TestEvery(t *testing.T) {
	t.Parallel()

	start := time.Now()
	var end time.Time

	ctx, cancel := context.WithCancel(context.Background())

	err := cron.Every(ctx, time.Second, func(ctx context.Context) error {
		end = time.Now()
		cancel()
		return nil
	})

	assert.True(t, errors.Is(err, context.Canceled))
	assert.True(t, end.Sub(start) >= time.Second && end.Sub(start) <= time.Second*2)
}

func TestAt(t *testing.T) {
	t.Parallel()

	start := time.Now().Add(time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	err := cron.At(ctx, start, func(ctx context.Context) error {
		cancel()
		return nil
	})

	assert.True(t, errors.Is(err, context.Canceled))
}
