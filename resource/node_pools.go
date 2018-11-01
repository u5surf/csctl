package resource

import (
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/table"
)

type nodePools struct {
	resource
	items []types.NodePool
}

func NewNodePools(items []types.NodePool) *nodePools {
	return &nodePools{
		resource: resource{
			name:    "nodepool",
			plural:  "nodepools",
			aliases: []string{"np", "nps"},
		},
		items: items,
	}
}

func NodePools() *nodePools {
	return NewNodePools(nil)
}

func (p *nodePools) columns() []string {
	return []string{
		"Name",
		"ID",
		"Mode",
		"Kubernetes Version",
		"Etcd Version",
		"Docker Version",
	}
}

func (p *nodePools) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, np := range p.items {
		etcdVersion := *np.EtcdVersion
		if etcdVersion == "" {
			etcdVersion = "N/A"
		}

		table.Append([]string{
			*np.Name,
			string(np.ID),
			*np.KubernetesMode,
			*np.KubernetesVersion,
			etcdVersion,
			*np.DockerVersion,
		})
	}

	table.Render()

	return nil
}

func (p *nodePools) JSON(w io.Writer) error {
	return displayJSON(w, p.items)
}

func (p *nodePools) YAML(w io.Writer) error {
	return displayYAML(w, p.items)
}

func (p *nodePools) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
