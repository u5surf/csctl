package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/resource/table"
)

// PluginsCatalog is a list of the associated cloud resource with additional functionality
type PluginsCatalog struct {
	resource
	items *types.PluginCatalog
}

// NewPluginCatalog constructs a new PluginsCatalog wrapping the given cloud type
func NewPluginCatalog(item *types.PluginCatalog) *PluginsCatalog {
	return &PluginsCatalog{
		resource: resource{
			name:     "plugin-catalog",
			aliases:  []string{"plgnc", "plugc", "pc"},
			listView: true,
		},
		items: item,
	}
}

// PluginCatalog constructs a new PluginCatalog with no underlying items, useful for
// interacting with the metadata itself.
func PluginCatalog() *PluginsCatalog {
	return NewPluginCatalog(nil)
}

// NewPluginCatalogFromDefinition constructs a new PluginsCatalog wrapping only
// the given PluginDefinition
func NewPluginCatalogFromDefinition(pluginType string, def *types.PluginDefinition) *PluginsCatalog {
	return NewPluginCatalogFromDefinitions(pluginType, []*types.PluginDefinition{def})
}

// NewPluginCatalogFromDefinitions constructs a new PluginsCatalog wrapping only
// the given PluginDefinitions
func NewPluginCatalogFromDefinitions(pluginType string, defs []*types.PluginDefinition) *PluginsCatalog {
	pc := types.PluginCatalog{}

	switch pluginType {
	case "autoscaler":
		pc.Autoscaler = defs
	case "cni":
		pc.CNI = defs
	case "csi":
		pc.CSI = defs
	case "cloud_controller_manager":
		pc.CloudControllerManager = defs
	case "cluster_management":
		pc.ClusterManagement = defs
	case "logs":
		pc.Logs = defs
	case "metrics":
		pc.Metrics = defs
	}

	return NewPluginCatalog(&pc)
}

// NewPluginCatalogFromVersion constructs a new PluginsCatalog wrapping only
// the given PluginVersion
func NewPluginCatalogFromVersion(pluginType, pluginImplementation string,
	version *types.PluginVersion) *PluginsCatalog {
	def := types.PluginDefinition{
		Type:           &pluginType,
		Implementation: &pluginImplementation,
		Versions:       []*types.PluginVersion{version},
	}

	return NewPluginCatalogFromDefinition(pluginType, &def)
}

func (pc *PluginsCatalog) columns() []string {
	return []string{
		"Type",
		"Implementation",
		"Versions",
		"Kubernetes Compatibility",
	}
}

// addPlugins adds plugins to the given table. It may be called with nil plugins,
// in which case it does nothing.
func addPlugins(plugins []*types.PluginDefinition, table *table.Table, ptype string) {
	if plugins == nil {
		return
	}

	for _, plugin := range plugins {
		for _, versionInfo := range plugin.Versions {
			table.Append([]string{
				ptype,
				*plugin.Implementation,
				*versionInfo.Version,
				*versionInfo.Compatibility.Kubernetes.Min + " - " + *versionInfo.Compatibility.Kubernetes.Max,
			})
		}
	}
}

// Table outputs the table representation to the given writer
func (pc *PluginsCatalog) Table(w io.Writer) error {
	table := table.New(w, pc.columns())

	addPlugins(pc.items.Autoscaler, table, "Autoscaler")
	addPlugins(pc.items.CloudControllerManager, table, "Cloud Controller Manager")
	addPlugins(pc.items.ClusterManagement, table, "Cluster Management")
	addPlugins(pc.items.CNI, table, "CNI")
	addPlugins(pc.items.CSI, table, "CSI")
	addPlugins(pc.items.Logs, table, "Logs")
	addPlugins(pc.items.Metrics, table, "Metrics")

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (pc *PluginsCatalog) JSON(w io.Writer) error {
	return displayJSON(w, pc.items, pc.resource.listView)
}

// YAML outputs the YAML representation to the given writer
func (pc *PluginsCatalog) YAML(w io.Writer) error {
	return displayYAML(w, pc.items, pc.resource.listView)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (pc *PluginsCatalog) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, pc.items)
}
