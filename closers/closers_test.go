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
	c := &NoopCloser{err: io.EOF}

	closers.Close(c)
	assert.True(t, c.closed)
}
