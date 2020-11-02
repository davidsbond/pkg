package cluster_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"pkg.dsb.dev/cluster"
)

func TestJoinLocal(t *testing.T) {
	node, err := cluster.JoinLocal()
	assert.NoError(t, err)
	assert.NotNil(t, node)

	peers := node.Peers()
	assert.Empty(t, peers)
	assert.NoError(t, node.Close())
}

func TestJoinLAN(t *testing.T) {
	node, err := cluster.JoinLAN()
	assert.NoError(t, err)
	assert.NotNil(t, node)

	peers := node.Peers()
	assert.Empty(t, peers)
	assert.NoError(t, node.Close())
}

func TestJoinWAN(t *testing.T) {
	node, err := cluster.JoinWAN()
	assert.NoError(t, err)
	assert.NotNil(t, node)

	peers := node.Peers()
	assert.Empty(t, peers)
	assert.NoError(t, node.Close())
}
