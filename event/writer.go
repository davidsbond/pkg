package event

import (
	"context"
	"net/url"
	"time"

	"github.com/opentracing/opentracing-go"
	"gocloud.dev/pubsub"

	"pkg.dsb.dev/tracing"
)

type (
	// The Writer type is used to write events to a single topic.
	Writer struct {
		topic *pubsub.Topic
		name  string
	}
)

// NewWriter creates a new instance of the Writer type that will write events to the configured
// event stream provider identified using the given URL.
func NewWriter(ctx context.Context, urlStr string) (*Writer, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	topic, err := pubsub.OpenTopic(ctx, urlStr)
	return &Writer{topic: topic, name: u.Host}, err
}

// Write an event to the stream.
func (w *Writer) Write(ctx context.Context, evt Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "event-write")
	defer span.Finish()

	span.SetTag("event.topic", w.name)

	err := w.topic.Send(ctx, &pubsub.Message{
		Body: evt.Payload,
	})
	if err != nil {
		return tracing.WithError(span, err)
	}

	eventsWritten.WithLabelValues(w.name).Inc()
	return nil
}

// Close the connection to the event stream.
func (w *Writer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return w.topic.Shutdown(ctx)
}
