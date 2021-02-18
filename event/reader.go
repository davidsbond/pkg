package event

import (
	"context"
	"time"

	"gocloud.dev/pubsub"

	"pkg.dsb.dev/tracing"
)

type (
	// The Reader type is used to handle inbound events from a single topic.
	Reader struct {
		subscription *pubsub.Subscription
	}
)

// NewReader creates a new instance of the Reader type that will read events from the configured
// event stream provider identified using the given URL.
func NewReader(ctx context.Context, urlStr string) (*Reader, error) {
	subscription, err := pubsub.OpenSubscription(ctx, urlStr)
	return &Reader{subscription: subscription}, err
}

// Read events from the stream, invoking fn for each inbound event. This method will block until fn returns
// an error or the provided context is cancelled.
func (r *Reader) Read(ctx context.Context, fn Handler) error {
	for ctx.Err() == nil {
		msg, err := r.subscription.Receive(ctx)
		if err != nil {
			return err
		}

		evt, err := unmarshal(msg.Body)
		if err != nil {
			return err
		}

		// If the message contains tracing information, start a new span as the child. This means traces work
		// across events.
		span, ctx, err := tracing.SpanFromMetadata(ctx, "event-read", evt.Sender.Metadata)
		if err != nil {
			return err
		}

		span.SetTag("event.type", evt.typeName())
		if err = fn(ctx, evt); err != nil {
			msg.Nack()
			err = tracing.WithError(span, err)
			span.Finish()
			return err
		}

		eventsRead.WithLabelValues(evt.typeName()).Inc()
		msg.Ack()
		span.Finish()
	}

	return ctx.Err()
}

// Close the connection to the event stream.
func (r *Reader) Close() error {
	const timeout = time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return r.subscription.Shutdown(ctx)
}
