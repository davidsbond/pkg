// Package cluster provides basic tools to build a self-aware cluster of applications. You can join a cluster and
// query peers that serve general build metadata.
package cluster

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/hashicorp/memberlist"

	"pkg.dsb.dev/environment"
	"pkg.dsb.dev/health"
	"pkg.dsb.dev/multierror"
)

type (
	// The Node type represents the local node in the cluster.
	Node struct {
		cluster *memberlist.Memberlist
	}

	// The Peer type represents a peer node in the cluster.
	Peer struct {
		node *memberlist.Node
	}

	// The Metadata type contains fields advertised by nodes in the cluster that describe
	// themselves.
	Metadata struct {
		Version                string    `json:"version,omitempty"`
		ApplicationName        string    `json:"application_name,omitempty"`
		ApplicationDescription string    `json:"application_description,omitempty"`
		Compiled               time.Time `json:"compiled,omitempty"`
	}
)

// JoinLAN attempts to join a LAN-based cluster. It uses the hostname as the node name,
// and otherwise sets very conservative values that are sane for most LAN environments.
func JoinLAN(opts ...Option) (*Node, error) {
	cnf := memberlist.DefaultLANConfig()

	return join(cnf, opts...)
}

// JoinLocal attempts to join a local cluster. It works similarly to JoinLAN, however it uses
// a configuration that is optimized for a local loopback environments.
func JoinLocal(opts ...Option) (*Node, error) {
	cnf := memberlist.DefaultLocalConfig()

	return join(cnf, opts...)
}

// JoinWAN attempts to join a WAN-based cluster. It uses a configuration that is optimized
// for most WAN environments.
func JoinWAN(opts ...Option) (*Node, error) {
	cnf := memberlist.DefaultWANConfig()

	return join(cnf, opts...)
}

func join(cnf *memberlist.Config, opts ...Option) (*Node, error) {
	cnf.Events = &metricsDelegate{}
	cnf.LogOutput = ioutil.Discard

	c := &config{ml: cnf}
	for _, opt := range opts {
		opt(c)
	}

	md, err := json.Marshal(Metadata{
		Version:                environment.Version,
		ApplicationName:        environment.ApplicationName,
		Compiled:               environment.Compiled(),
		ApplicationDescription: environment.ApplicationDescription,
	})
	if err != nil {
		return nil, err
	}

	cluster, err := memberlist.Create(cnf)
	if err != nil {
		return nil, err
	}

	cluster.LocalNode().Meta = md

	if _, err = cluster.Join(c.nodes); err != nil {
		return nil, err
	}

	node := &Node{cluster: cluster}
	health.AddCheck("cluster", node.Ping)
	return node, nil
}

// Peers returns all currently known Peers in the cluster.
func (n *Node) Peers() (peers []*Peer) {
	for _, member := range n.cluster.Members() {
		if member == n.cluster.LocalNode() {
			continue
		}

		peers = append(peers, &Peer{node: member})
	}

	return
}

// Close the connection to the cluster.
func (n *Node) Close() error {
	return multierror.Append(
		n.cluster.Leave(time.Minute),
		n.cluster.Shutdown(),
	)
}

// Ping returns a non-nil error if the node is not functioning as expected.
func (n *Node) Ping() error {
	return n.cluster.UpdateNode(time.Minute)
}

// Address returns the address of the peer, without a port.
func (p *Peer) Address() string {
	return p.node.Addr.String()
}

// Name returns the advertised name of the peer.
func (p *Peer) Name() string {
	return p.node.Name
}

// Metadata unmarshals the peer's advertised metadata and returns it.
func (p *Peer) Metadata() (*Metadata, error) {
	var out Metadata
	err := json.Unmarshal(p.node.Meta, &out)

	return &out, err
}
