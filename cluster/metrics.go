package cluster

import (
	"github.com/hashicorp/memberlist"
	"github.com/prometheus/client_golang/prometheus"

	"pkg.dsb.dev/metrics"
)

const (
	namespace = "cluster"
	subsystem = "event"
)

func init() {
	metrics.Register(nodesJoined, nodesLeft, nodesUpdated)
}

var (
	nodesJoined = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "node_joined",
		Help:      "Count of node join events received",
	})

	nodesUpdated = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "node_updated",
		Help:      "Count of node update events received",
	})

	nodesLeft = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "node_leave",
		Help:      "Count of node leave events received",
	})
)

type (
	metricsDelegate struct{}
)

func (m *metricsDelegate) NotifyJoin(_ *memberlist.Node) {
	nodesJoined.Inc()
}

func (m *metricsDelegate) NotifyLeave(_ *memberlist.Node) {
	nodesLeft.Inc()
}

func (m *metricsDelegate) NotifyUpdate(_ *memberlist.Node) {
	nodesUpdated.Inc()
}
