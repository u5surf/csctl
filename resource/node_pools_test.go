package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	np1     = "test1"
	np2     = "test2"
	master  = "master"
	worker  = "worker"
	version = "v1.0.0"

	nps = []types.NodePool{
		{
			Name:              &np1,
			ID:                types.UUID("1234"),
			KubernetesMode:    &master,
			KubernetesVersion: &version,
			EtcdVersion:       &version,
			DockerVersion:     &version,
		},
		{
			Name:              &np2,
			ID:                types.UUID("4321"),
			KubernetesMode:    &worker,
			KubernetesVersion: &version,
			EtcdVersion:       nil,
			DockerVersion:     &version,
		},
	}
)

func TestNewNodePools(t *testing.T) {
	a := NewNodePools(nil)
	assert.NotNil(t, a)

	a = NewNodePools(nps)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(nps))

	a = NodePool()
	assert.NotNil(t, a)
}

func TestNodePoolsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewNodePools(nps)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(nps), info.numRows)
}
