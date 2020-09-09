// Package event contains utilities for interacting with various event-stream providers. Including the ability
// to write and read from event-streaming sources.
package event

import (
	// Gocloud driver for kafka event streaming.
	_ "gocloud.dev/pubsub/kafkapubsub"
	// Gocloud driver for in-memory event streaming.
	_ "gocloud.dev/pubsub/mempubsub"
)

type (
	// The Event type represents an event that can be written or read from the stream. It just contains a slice
	// of bytes that the user must marshal into whatever types they desire. It does not assume serialization format, so
	// anything that can be marshalled into a slice of bytes can be used.
	Event struct {
		Topic   string // Set internally when reading, not used when writing.
		Payload []byte
	}
)
