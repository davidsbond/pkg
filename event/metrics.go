package event

import (
	"github.com/prometheus/client_golang/prometheus"

	"pkg.dsb.dev/metrics"
)

const (
	namespace = "event"
	subsystem = "stream"
)

func init() {
	metrics.Register(eventsWritten, eventsRead)
}

var (
	eventsWritten = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "events_written",
		Help:      "Total number of events written to the stream",
	}, []string{"type"})

	eventsRead = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "events_read",
		Help:      "Total number of events read from the stream",
	}, []string{"type"})
)
