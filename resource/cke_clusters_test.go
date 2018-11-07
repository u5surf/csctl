package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	ckeClusterTime = "1517001176920"

	ckeClusters = []types.CKECluster{
		{
			ID:           types.UUID("1234"),
			ProviderName: strptr("google"),
			Status: &types.CKEClusterStatus{
				Type: strptr("RUNNING"),
			},
			OwnerID:   types.UUID("1234"),
			CreatedAt: &ckeClusterTime,
		},
		{
			ID:           types.UUID("4321"),
			ProviderName: strptr("amazon_web_services"),
			Status: &types.CKEClusterStatus{
				Type: strptr("PROVISIONING"),
			},
			OwnerID:   types.UUID("4321"),
			CreatedAt: &ckeClusterTime,
		},
	}
)

func TestNewCKEClusters(t *testing.T) {
	a := NewCKEClusters(nil)
	assert.NotNil(t, a)

	a = NewCKEClusters(ckeClusters)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(ckeClusters))

	a = CKECluster()
	assert.NotNil(t, a)
}

func TestCKEClustersTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewCKEClusters(ckeClusters)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(ckeClusters), info.numRows)
}
