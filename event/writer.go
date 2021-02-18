package event

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"gocloud.dev/pubsub"

	"pkg.dsb.dev/tracing"
)

type (
	// The Writer type is used to write events to a single topic.
	Writer struct {
		topic *pubsub.Topic
	}
)

// NewWriter creates a new instance of the Writer type that will write events to the configured
// event stream provider identified using the given URL.
func NewWriter(ctx context.Context, urlStr string) (*Writer, error) {
	topic, err := pubsub.OpenTopic(ctx, urlStr)
	return &Writer{topic: topic}, err
}

// Write an event to the stream. If the provided context.Context contains tracing information, it is added to
// Event.Sender.Metadata so that tracing can occur across event readers/writers.
func (w *Writer) Write(ctx context.Context, evt Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "event-write")
	defer span.Finish()

	span.SetTag("event.type", evt.typeName())

	md, err := tracing.SpanMetadata(span)
	if err != nil {
		return err
	}
	for k, v := range md {
		evt.Sender.Metadata[k] = v
	}

	body, err := evt.marshal()
	if err != nil {
		return err
	}

	err = w.topic.Send(ctx, &pubsub.Message{
		Body: body,
	})

	if err != nil {
		return tracing.WithError(span, err)
	}

	eventsWritten.WithLabelValues(evt.typeName()).Inc()
	return nil
}

// Close the connection to the event stream.
func (w *Writer) Close() error {
	const timeout = time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return w.topic.Shutdown(ctx)
}
