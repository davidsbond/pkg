// Package queue contains a simple job queue implementation for running concurrent jobs.
package queue

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"

	"pkg.dsb.dev/tracing"
)

type (
	// The Job type represents a single unit of work.
	Job struct {
		ID string
	}

	// WorkFunc is a function invoked for each job.
	WorkFunc func(ctx context.Context, j Job) error

	// The Queue type is responsible for queueing and executing concurrent jobs.
	Queue struct {
		errors     chan error
		maxWorkers int
		queue      chan Job
		wf         WorkFunc
	}
)

// New creates a new *Queue instance with the specified configuration options.
func New(ctx context.Context, wf WorkFunc, opts ...Option) *Queue {
	c := config{
		maxWorkers: 1,
	}

	for _, opt := range opts {
		opt(&c)
	}

	q := &Queue{
		errors:     make(chan error),
		maxWorkers: c.maxWorkers,
		queue:      make(chan Job),
		wf:         wf,
	}

	go q.start(ctx)
	return q
}

// Work the queue until there are no jobs left.
func (q *Queue) start(ctx context.Context) {
	grp, ctx := errgroup.WithContext(ctx)

	for i := 0; i < q.maxWorkers; i++ {
		grp.Go(func() error {
			return q.work(ctx, q.wf)
		})
	}

	grp.Go(func() error {
		<-ctx.Done()
		close(q.queue)
		close(q.errors)
		return nil
	})

	err := grp.Wait()
	if span := opentracing.SpanFromContext(ctx); span != nil && err != nil {
		q.errors <- tracing.WithError(span, err)
	}
}

func (q *Queue) work(ctx context.Context, wf WorkFunc) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case job, ok := <-q.queue:
			if !ok {
				return nil
			}

			if err := wf(ctx, job); err != nil {
				return err
			}
		}
	}
}

// Add a job to the queue.
func (q *Queue) Add(ctx context.Context, job Job) error {
	select {
	case q.queue <- job:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Errors exposes the errors channel for reading.
func (q *Queue) Errors() <-chan error {
	return q.errors
}
