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

// NewPluginCatalog constructs a new PluginCatalog wrapping the given cloud type
func NewPluginCatalog(item *types.PluginCatalog) *PluginsCatalog {
	return &PluginsCatalog{
		resource: resource{
			name:    "plugin-catalog",
			aliases: []string{"plgnc", "plugc", "pc"},
		},
		items: item,
	}
}

// PluginCatalog constructs a new PluginCatalog with no underlying items, useful for
// interacting with the metadata itself.
func PluginCatalog() *PluginsCatalog {
	return NewPluginCatalog(nil)
}

func (pc *PluginsCatalog) columns() []string {
	return []string{
		"Type",
		"Implementation",
		"Versions",
		"Kubernetes Compatibility",
	}
}

func addPlugins(plugins []*types.PluginDefinition, table *table.Table, ptype string) {
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

	addPlugins(pc.items.CloudControllerManager, table, "Cloud Controller Manger")
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
	return displayJSON(w, pc.items)
}

// YAML outputs the YAML representation to the given writer
func (pc *PluginsCatalog) YAML(w io.Writer) error {
	return displayYAML(w, pc.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (pc *PluginsCatalog) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, pc.items)
}
