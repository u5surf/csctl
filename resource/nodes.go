package resource

import (
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Nodes is a list of the associated cloud resource with additional functionality
type Nodes struct {
	resource
	items []types.Node
}

// NewNodes constructs a new Nodes wrapping the given cloud type
func NewNodes(items []types.Node) *Nodes {
	return &Nodes{
		resource: resource{
			name:     "node",
			plural:   "nodes",
			aliases:  []string{"no", "nos"},
			listView: true,
		},
		items: items,
	}
}

// Node constructs a new Nodes with no underlying items, useful for
// interacting with the metadata itself.
func Node() *Nodes {
	return NewNodes(nil)
}

func (p *Nodes) columns() []string {
	return []string{
		"ID",
		"Status",
		"Created At",
		"Updated At",
	}
}

// Table outputs the table representation to the given writer
func (p *Nodes) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, node := range p.items {
		var status string
		if node.Status == nil || node.Status.Type == nil ||
			*node.Status.Type == "" {
			status = "UNKNOWN"
		} else {
			status = *node.Status.Type
		}

		table.Append([]string{
			string(node.ID),
			status,
			convert.UnixTimeMSToString(*node.CreatedAt),
			convert.UnixTimeMSToString(*node.UpdatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *Nodes) JSON(w io.Writer) error {
	return displayJSON(w, p.items, p.resource.listView)
}

// YAML outputs the YAML representation to the given writer
func (p *Nodes) YAML(w io.Writer) error {
	return displayYAML(w, p.items, p.resource.listView)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *Nodes) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
