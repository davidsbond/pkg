// Package event contains utilities for interacting with various event-stream providers. Including the ability
// to write and read from event-streaming sources.
package event

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"pkg.dsb.dev/environment"
	"pkg.dsb.dev/event/v1"

	// Gocloud driver for kafka event streaming.
	_ "gocloud.dev/pubsub/kafkapubsub"
	// Gocloud driver for gcp event streaming.
	_ "gocloud.dev/pubsub/gcppubsub"
	// Gocloud driver for in-memory event streaming.
	_ "gocloud.dev/pubsub/mempubsub"
	// Gocloud driver for nats event streaming.
	_ "gocloud.dev/pubsub/natspubsub"
	// Gocloud driver for rabbitmq event streaming.
	_ "gocloud.dev/pubsub/rabbitpubsub"
	// Gocloud driver for azure event streaming.
	_ "gocloud.dev/pubsub/azuresb"
	// Gocloud driver for aws event streaming.
	_ "gocloud.dev/pubsub/awssnssqs"
)

type (
	// The Event type describes something that has happened at a particular point in time. It contains information
	// on who sent it and when. Each event payload is a proto-encoded message.
	Event struct {
		ID        string
		Timestamp time.Time
		AppliesAt time.Time
		Payload   proto.Message
		Sender    Sender
	}

	// Sender contains details on who sent an Event.
	Sender struct {
		Application string
		Metadata    map[string]string
	}

	// The Handler type is a function that processes an inbound event.
	Handler func(ctx context.Context, evt Event) error
)

// New returns a new Event instance that contains the provided proto.Message implementation. A unique identifier
// is generated, timestamps are set to now and the application name is taken from environment.ApplicationName.
func New(msg proto.Message) Event {
	return Event{
		ID:        uuid.Must(uuid.NewV4()).String(),
		Timestamp: time.Now(),
		AppliesAt: time.Now(),
		Payload:   msg,
		Sender: Sender{
			Application: environment.ApplicationName,
			Metadata:    make(map[string]string),
		},
	}
}

func (e Event) typeName() string {
	return string(e.Payload.ProtoReflect().Descriptor().FullName())
}

func (e Event) marshal() ([]byte, error) {
	any, err := anypb.New(e.Payload)
	if err != nil {
		return nil, err
	}

	envelope := &event.Envelope{
		Id:        e.ID,
		Timestamp: timestamppb.New(e.Timestamp),
		AppliesAt: timestamppb.New(e.AppliesAt),
		Payload:   any,
		Sender: &event.Sender{
			Application: e.Sender.Application,
			Metadata:    e.Sender.Metadata,
		},
	}

	return proto.Marshal(envelope)
}

func unmarshal(b []byte) (Event, error) {
	var env event.Envelope
	if err := proto.Unmarshal(b, &env); err != nil {
		return Event{}, err
	}

	payload, err := env.Payload.UnmarshalNew()
	if err != nil {
		return Event{}, err
	}

	return Event{
		ID:        env.GetId(),
		Timestamp: env.GetTimestamp().AsTime(),
		AppliesAt: env.GetAppliesAt().AsTime(),
		Payload:   payload,
		Sender: Sender{
			Application: env.GetSender().GetApplication(),
			Metadata:    env.GetSender().GetMetadata(),
		},
	}, nil
}
