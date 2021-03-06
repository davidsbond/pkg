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
	"google.golang.org/protobuf/types/known/timestamppb"

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

	payload := timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	expected := event.New(payload)

	assert.NoError(t, wr.Write(ctx, expected))
	err := rd.Read(ctx, func(ctx context.Context, actual event.Event) error {
		assert.EqualValues(t, expected.ID, actual.ID)
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

	payload := timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	expected := event.New(payload)

	assert.NoError(t, wr.Write(ctx, expected))

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
