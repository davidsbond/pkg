package database_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/storage/database"
)

func TestFromInterval(t *testing.T) {
	t.Parallel()

	exp := time.Hour
	itv := database.ToInterval(exp)
	act := database.FromInterval(itv)

	assert.Equal(t, exp, act)
}

func TestFromTextArray(t *testing.T) {
	t.Parallel()

	exp := []string{"a", "b", "c"}
	arr := database.ToTextArray(exp)
	act := database.FromTextArray(arr)

	assert.EqualValues(t, exp, act)
}
