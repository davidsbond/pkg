package event_test

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-client-go"

	"pkg.dsb.dev/closers"
	"pkg.dsb.dev/event"
)

var (
	wr *event.Writer
	rd *event.Reader
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	var err error

	wr, err = event.NewWriter(ctx, "mem://test-topic")
	if err != nil {
		log.Fatal(err)
	}
	defer closers.Close(wr)

	rd, err = event.NewReader(ctx, "mem://test-topic")
	if err != nil {
		log.Fatal(err)
	}
	defer closers.Close(rd)

	os.Exit(m.Run())
}

func TestReader_Read(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	assert.NoError(t, wr.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))

	err := rd.Read(ctx, func(ctx context.Context, evt event.Event) error {
		assert.EqualValues(t, []byte("hello world"), evt.Payload)
		cancel()
		return nil
	})
	assert.True(t, errors.Is(err, context.Canceled))
}

func TestEventTracing(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	tracer, closer := jaeger.NewTracer(
		"test",
		jaeger.NewConstSampler(true),
		jaeger.NewInMemoryReporter(),
	)

	defer closers.Close(closer)
	opentracing.SetGlobalTracer(tracer)

	assert.NoError(t, wr.Write(ctx, event.Event{
		Payload: []byte("hello world"),
	}))

	err := rd.Read(ctx, func(ctx context.Context, evt event.Event) error {
		span := opentracing.SpanFromContext(ctx)
		assert.NotNil(t, span)

		spanCtx := span.Context()
		assert.NotNil(t, spanCtx)

		cancel()
		return nil
	})
	assert.True(t, errors.Is(err, context.Canceled))
}
