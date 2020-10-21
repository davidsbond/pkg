package closers_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/closers"
)

type (
	NoopCloser struct {
		err    error
		closed bool
	}
)

func (c *NoopCloser) Close() error {
	c.closed = true
	return c.err
}

func TestClose(t *testing.T) {
	t.Parallel()

	c := &NoopCloser{err: io.EOF}

	closers.Close(c)
	assert.True(t, c.closed)
}

func TestCloseFunc(t *testing.T) {
	t.Parallel()

	called := false
	c := closers.CloseFunc(func() error {
		called = true
		return nil
	})

	closers.Close(c)
	assert.True(t, called)
}
