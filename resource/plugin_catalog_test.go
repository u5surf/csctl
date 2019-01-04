package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	plugsCatalog = &types.PluginCatalog{
		Autoscaler: []*types.PluginDefinition{
			{
				Implementation: strptr("cerebral"),
				Versions: []*types.PluginVersion{
					{
						Version: strptr("1.0.0"),
						Compatibility: &types.PluginCompatibility{
							Kubernetes: &types.PluginKubernetesCompatibility{
								Min: strptr("1.10.x"),
								Max: strptr("1.12.x"),
							},
							Plugins: map[string]types.PluginPluginsCompatibility{
								"Metrics": {
									Implementation: strptr("prometheus"),
									Min:            strptr("1.0.0"),
									Max:            strptr("1.1.0"),
								},
							},
						},
					},
				},
			},
		},
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

func TestNewPluginCatalogFromDefinition(t *testing.T) {
	pc := NewPluginCatalogFromDefinition("cni", nil)
	assert.NotNil(t, pc, "nil definition is ok; ,result is empty catalog")

	pc = NewPluginCatalogFromDefinition("cni", plugsCatalog.CNI[0])
	assert.Equal(t, plugsCatalog.CNI[0], pc.items.CNI[0], "cni def set properly")

	// We'll only do these checks once. Doesn't seem worth the effort to do it
	// for every type, as the code is straightforward.
	assert.Nil(t, pc.items.Autoscaler, "only cni def is set")
	assert.Nil(t, pc.items.CSI, "only cni def is set")
	assert.Nil(t, pc.items.CloudControllerManager, "only cni def is set")
	assert.Nil(t, pc.items.ClusterManagement, "only cni def is set")
	assert.Nil(t, pc.items.Logs, "only cni def is set")
	assert.Nil(t, pc.items.Metrics, "only cni def is set")

	pc = NewPluginCatalogFromDefinition("autoscaler", plugsCatalog.Autoscaler[0])
	assert.Equal(t, plugsCatalog.Autoscaler[0], pc.items.Autoscaler[0], "autoscaler def set properly")

	pc = NewPluginCatalogFromDefinition("csi", plugsCatalog.CSI[0])
	assert.Equal(t, plugsCatalog.CSI[0], pc.items.CSI[0], "csi def set properly")

	pc = NewPluginCatalogFromDefinition("cloud_controller_manager", plugsCatalog.CloudControllerManager[0])
	assert.Equal(t, plugsCatalog.CloudControllerManager[0], pc.items.CloudControllerManager[0],
		"ccm def set properly")

	pc = NewPluginCatalogFromDefinition("cluster_management", plugsCatalog.ClusterManagement[0])
	assert.Equal(t, plugsCatalog.ClusterManagement[0], pc.items.ClusterManagement[0],
		"cluster_management def set properly")

	pc = NewPluginCatalogFromDefinition("logs", plugsCatalog.Logs[0])
	assert.Equal(t, plugsCatalog.Logs[0], pc.items.Logs[0], "logs def set properly")

	pc = NewPluginCatalogFromDefinition("metrics", plugsCatalog.Metrics[0])
	assert.Equal(t, plugsCatalog.Metrics[0], pc.items.Metrics[0], "metrics def set properly")
}

func TestNewPluginCatalogFromVersion(t *testing.T) {
	pc := NewPluginCatalogFromVersion("cni", "calico", nil)
	assert.NotNil(t, pc, "nil version is ok; ,result is empty catalog")

	pc = NewPluginCatalogFromVersion("cni", "calico", plugsCatalog.CNI[0].Versions[0])
	assert.Equal(t, plugsCatalog.CNI[0].Versions[0], pc.items.CNI[0].Versions[0], "cni version set properly")
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
	l += len(plugsCatalog.Autoscaler)
	l += len(plugsCatalog.CNI)
	l += len(plugsCatalog.CSI)
	l += len(plugsCatalog.CloudControllerManager)
	l += len(plugsCatalog.ClusterManagement)
	l += len(plugsCatalog.Logs)
	l += len(plugsCatalog.Metrics)

	assert.Equal(t, l, info.numRows)
}
