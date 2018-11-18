package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	plugsCatalog = &types.PluginCatalog{
		CNI: []*types.PluginDefinition{
			{
				Implementation: strptr("calico"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("2.0.0"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.10.x"),
								Max: strptr("1.12.x"),
							},
						},
					},
				},
			},
		},
		CSI: []*types.PluginDefinition{
			{
				Implementation: strptr("digitalocean"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("0.3.0"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.12.x"),
								Max: strptr("1.12.x"),
							},
						},
					},
				},
			},
		},
		CloudControllerManager: []*types.PluginDefinition{
			{
				Implementation: strptr("digitalocean"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("1.0.0"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.10.x"),
								Max: strptr("1.12.x"),
							},
						},
					},
				},
			},
		},
		ClusterManagement: []*types.PluginDefinition{
			{
				Implementation: strptr("containership"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("4.0.1"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.10.x"),
								Max: strptr("1.12.x"),
							},
						},
					},
				},
			},
		},
		Logs: []*types.PluginDefinition{
			{
				Implementation: strptr("kubernetes"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("1.0.0"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.10.x"),
								Max: strptr("1.12.x"),
							},
						},
					},
				},
			},
		},
		Metrics: []*types.PluginDefinition{
			{
				Implementation: strptr("prometheus"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("1.0.0"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.10.x"),
								Max: strptr("1.12.x"),
							},
						},
					},
				},
			},
		},
	}
)

func TestNewPluginCatalog(t *testing.T) {
	pc := NewPluginCatalog(nil)
	assert.NotNil(t, pc)

	pc = NewPluginCatalog(plugsCatalog)
	assert.NotNil(t, pc)
	assert.Equal(t, pc.items, plugsCatalog)

	pc = PluginCatalog()
	assert.NotNil(t, pc)
}

func TestPluginCatalogTable(t *testing.T) {
	buf := new(bytes.Buffer)

	pc := NewPluginCatalog(plugsCatalog)
	assert.NotNil(t, pc)

	err := pc.Table(buf)
	assert.NoError(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(pc.columns()), info.numHeaderCols)
	assert.Equal(t, len(pc.columns()), info.numCols)

	l := 0
	l += len(plugsCatalog.CNI)
	l += len(plugsCatalog.CSI)
	l += len(plugsCatalog.CloudControllerManager)
	l += len(plugsCatalog.ClusterManagement)
	l += len(plugsCatalog.Logs)
	l += len(plugsCatalog.Metrics)

	assert.Equal(t, l, info.numRows)
}
