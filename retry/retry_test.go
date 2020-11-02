package retry_test

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/retry"
)

func TestDo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Test no error only invokes once
	invoked := 0
	assert.NoError(t, retry.Do(ctx, 3, func(_ context.Context) error {
		invoked++
		return nil
	}))

	assert.EqualValues(t, 1, invoked)

	// Test error invokes the maximum times before returning
	invoked = 0
	assert.Error(t, retry.Do(ctx, 3, func(_ context.Context) error {
		invoked++
		return io.EOF
	}))

	assert.EqualValues(t, 4, invoked)
}

func TestStop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	invoked := 0
	assert.Error(t, retry.Do(ctx, 3, func(_ context.Context) error {
		invoked++
		return retry.Stop(io.EOF)
	}))

	assert.EqualValues(t, 1, invoked)
}
