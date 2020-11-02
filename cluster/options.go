package cluster

import (
	"log"

	"github.com/hashicorp/memberlist"
)

type (
	// The Option type is a function that modifies the cluster configuration.
	Option func(cnf *config)

	config struct {
		ml    *memberlist.Config
		nodes []string
	}
)

// WithName sets the advertised name of the node, this should be unique across the cluster.
func WithName(name string) Option {
	return func(cnf *config) {
		cnf.ml.Name = name
	}
}

// WithSecretKey sets the secret key to use to encrypt messages between nodes.
func WithSecretKey(key string) Option {
	return func(cnf *config) {
		cnf.ml.SecretKey = []byte(key)
	}
}

// WithNodes sets a slice of addresses that the local node will connect to when the cluster
// is joined.
func WithNodes(nodes []string) Option {
	return func(cnf *config) {
		cnf.nodes = nodes
	}
}

// WithLogger sets the logger to be used by the discovery mechanism. By default, logging is
// disabled.
func WithLogger(l *log.Logger) Option {
	return func(cnf *config) {
		cnf.ml.LogOutput = nil
		cnf.ml.Logger = l
	}
}
