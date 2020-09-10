package distance_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/distance"
)

func TestBetween(t *testing.T) {
	t.Parallel()

	const (
		lat1 = -5.21702
		lon1 = -154.49308
		lat2 = 45.78295
		lon2 = 155.29529

		expected = 7529.1614370609595
	)

	actual := distance.Between(lat1, lon1, lat2, lon2)
	assert.EqualValues(t, expected, actual)
}
