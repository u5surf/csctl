package resource

import (
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// CKEClusters is a list of the associated cloud resource with additional functionality
type CKEClusters struct {
	resource
	filterable
	items []types.CKECluster
}

// NewCKEClusters constructs a new CKEClusters wrapping the given cloud type
func NewCKEClusters(items []types.CKECluster) *CKEClusters {
	return &CKEClusters{
		resource: resource{
			name:   "cluster",
			plural: "clusters",
		},
		items: items,
	}
}

// CKECluster constructs a new CKEClusters with no underlying items, useful for
// interacting with the metadata itself.
func CKECluster() *CKEClusters {
	return NewCKEClusters(nil)
}

func (c *CKEClusters) columns() []string {
	return []string{
		"ID",
		"Provider Name",
		"Status",
		"Owner ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (c *CKEClusters) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, cluster := range c.items {
		table.Append([]string{
			string(cluster.ID),
			*cluster.ProviderName,
			*cluster.Status.Type,
			string(cluster.OwnerID),
			convert.UnixTimeMSToString(*cluster.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *CKEClusters) JSON(w io.Writer) error {
	return displayJSON(w, c.items)
}

// YAML outputs the YAML representation to the given writer
func (c *CKEClusters) YAML(w io.Writer) error {
	return displayYAML(w, c.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *CKEClusters) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}

// FilterByOwnerID filters the underlying items by owner ID
func (c *CKEClusters) FilterByOwnerID(id string) {
	filtered := make([]types.CKECluster, 0)
	for _, cluster := range c.items {
		if string(cluster.OwnerID) == id {
			filtered = append(filtered, cluster)
		}
	}

	c.items = filtered
}
