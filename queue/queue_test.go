package queue_test

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/environment"
	"pkg.dsb.dev/queue"
)

func TestQueue_Work(t *testing.T) {
	ctx, cancel := context.WithTimeout(environment.NewContext(), time.Second*2)
	defer cancel()

	actual := make([]queue.Job, 0)
	mu := &sync.Mutex{}
	q := queue.New(ctx, func(ctx context.Context, j queue.Job) error {
		mu.Lock()
		defer mu.Unlock()

		actual = append(actual, j)
		return nil
	})

	expected := make([]queue.Job, 100)
	for i := 0; i < 100; i++ {
		j := queue.Job{ID: strconv.Itoa(i)}
		expected[i] = j
		assert.NoError(t, q.Add(ctx, j))
	}

	<-ctx.Done()
	mu.Lock()
	defer mu.Unlock()
	assert.EqualValues(t, expected, actual)
}
