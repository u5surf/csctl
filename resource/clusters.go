package resource

import (
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

type ckeClusters struct {
	resource
	filterable
	items []types.CKECluster
}

func NewCKEClusters(items []types.CKECluster) *ckeClusters {
	return &ckeClusters{
		resource: resource{
			name:   "cluster",
			plural: "clusters",
		},
		items: items,
	}
}

func CKEClusters() *ckeClusters {
	return NewCKEClusters(nil)
}

func (c *ckeClusters) columns() []string {
	return []string{
		"ID",
		"Provider Name",
		"Status",
		"Owner ID",
		"Created At",
	}
}

func (c *ckeClusters) Table(w io.Writer) error {
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

func (c *ckeClusters) JSON(w io.Writer) error {
	return displayJSON(w, c.items)
}

func (c *ckeClusters) YAML(w io.Writer) error {
	return displayYAML(w, c.items)
}

func (c *ckeClusters) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}

func (c *ckeClusters) FilterByOwnerID(id string) {
	filtered := make([]types.CKECluster, 0)
	for _, cluster := range c.items {
		if string(cluster.OwnerID) == id {
			filtered = append(filtered, cluster)
		}
	}

	c.items = filtered
}
