package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	etcdVersion   = "3.2.24"
	dockerVersion = "18.6.1"

	nps = []types.NodePool{
		{
			Name:              strptr("test1"),
			ID:                types.UUID("1234"),
			KubernetesMode:    strptr("master"),
			KubernetesVersion: strptr("1.12.1"),
			EtcdVersion:       &etcdVersion,
			DockerVersion:     &dockerVersion,
		},
		{
			Name:              strptr("test2"),
			ID:                types.UUID("4321"),
			KubernetesMode:    strptr("worker"),
			KubernetesVersion: strptr("1.11.1"),
			EtcdVersion:       nil,
			DockerVersion:     &dockerVersion,
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
