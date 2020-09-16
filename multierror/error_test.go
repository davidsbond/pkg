package multierror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/multierror"
)

func TestAppend(t *testing.T) {
	t.Parallel()

	// If we pass in nil, we get nil back.
	multiErr := multierror.Append(nil, nil, nil)
	assert.Nil(t, multiErr)

	first := errors.New("first")
	second := errors.New("second")
	third := errors.New("third")

	// Create an error with the first two errors.
	multiErr = multierror.Append(first, second)

	// Assert the error equals the first and second
	assert.True(t, errors.Is(multiErr, first))
	assert.True(t, errors.Is(multiErr, second))
	assert.Equal(t, "first; second", multiErr.Error())

	// Add a third error
	multiErr = multierror.Append(multiErr, third)

	// Assert we still equal all other errors.
	assert.True(t, errors.Is(multiErr, first))
	assert.True(t, errors.Is(multiErr, second))
	assert.True(t, errors.Is(multiErr, third))
	assert.Equal(t, "first; second; third", multiErr.Error())

	// Assert we don't equal some other random error
	fourth := errors.New("fourth")
	assert.False(t, errors.Is(multiErr, fourth))
}
