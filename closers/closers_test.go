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

func TestCloseAll(t *testing.T) {
	t.Parallel()

	cs := make([]io.Closer, 100)
	noops := make([]*NoopCloser, 100)
	for i := 0; i < 100; i++ {
		noops[i] = &NoopCloser{}
		cs[i] = noops[i]
	}

	closers.CloseAll(cs...)
	for _, c := range noops {
		assert.True(t, c.closed)
	}
}

func TestMultiCloser_Close(t *testing.T) {
	t.Parallel()

	mc := closers.MultiCloser{}
	noops := make([]*NoopCloser, 100)
	for i := 0; i < 100; i++ {
		noops[i] = &NoopCloser{}
		mc.Add(noops[i])
	}

	assert.NoError(t, mc.Close())
	for _, c := range noops {
		assert.True(t, c.closed)
	}
}
