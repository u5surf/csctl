package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Registries is a list of the associated cloud resource with additional functionality
type Registries struct {
	resource
	filterable
	items []types.Registry
}

// NewRegistries constructs a new Registries wrapping the given cloud type
func NewRegistries(items []types.Registry) *Registries {
	return &Registries{
		resource: resource{
			name:    "registry",
			plural:  "registries",
			aliases: []string{"reg", "regs"},
		},
		items: items,
	}
}

// Registry constructs a new Registries with no underlying items, useful for
// interacting with the metadata itself.
func Registry() *Registries {
	return NewRegistries(nil)
}

func (p *Registries) columns() []string {
	return []string{
		"ID",
		"Description",
		"Provider Name",
		"Server Address",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (p *Registries) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, reg := range p.items {
		table.Append([]string{
			string(reg.ID),
			*reg.Description,
			*reg.Provider,
			*reg.Serveraddress,
			convert.UnixTimeMSToString(*reg.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *Registries) JSON(w io.Writer) error {
	return displayJSON(w, p.items)
}

// YAML outputs the YAML representation to the given writer
func (p *Registries) YAML(w io.Writer) error {
	return displayYAML(w, p.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *Registries) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
