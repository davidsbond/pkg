package storage_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/storage"
)

func TestFromInterval(t *testing.T) {
	t.Parallel()

	exp := time.Hour
	itv := storage.ToInterval(exp)
	act := storage.FromInterval(itv)

	assert.Equal(t, exp, act)
}

func TestFromTextArray(t *testing.T) {
	t.Parallel()

	exp := []string{"a", "b", "c"}
	arr := storage.ToTextArray(exp)
	act := storage.FromTextArray(arr)

	assert.EqualValues(t, exp, act)
}
