package random_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/random"
)

func TestInt(t *testing.T) {
	t.Parallel()

	const max = 10

	assert.NotPanics(t, func() {
		result := random.Int(max)
		assert.LessOrEqual(t, result, max)
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	var actual string
	assert.NotPanics(t, func() {
		actual = random.String(10)
	})

	assert.NotEmpty(t, actual)
	assert.Len(t, actual, 10)
}

func TestCapitalisedAlphaNumeric(t *testing.T) {
	t.Parallel()

	var actual string
	assert.NotPanics(t, func() {
		actual = random.CapitalisedAlphaNumeric(10)
	})

	assert.NotEmpty(t, actual)
	assert.Len(t, actual, 10)
	assert.Regexp(t, "^[A-Z0-9]*$", actual)
}
