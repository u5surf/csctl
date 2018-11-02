package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Plugins is a list of the associated cloud resource with additional functionality
type Plugins struct {
	resource
	filterable
	items []types.Plugin
}

// NewPlugins constructs a new Plugins wrapping the given cloud type
func NewPlugins(items []types.Plugin) *Plugins {
	return &Plugins{
		resource: resource{
			name:    "plugin",
			plural:  "Plugins",
			aliases: []string{"plug", "plugs", "plgn", "plgns"},
		},
		items: items,
	}
}

// Plugin constructs a new Plugins with no underlying items, useful for
// interacting with the metadata itself.
func Plugin() *Plugins {
	return NewPlugins(nil)
}

func (p *Plugins) columns() []string {
	return []string{
		"ID",
		"Type",
		"Implementation",
		"Version",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (p *Plugins) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, plug := range p.items {
		table.Append([]string{
			string(plug.ID),
			*plug.Type,
			*plug.Implementation,
			*plug.Version,
			convert.UnixTimeMSToString(*plug.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *Plugins) JSON(w io.Writer) error {
	return displayJSON(w, p.items)
}

// YAML outputs the YAML representation to the given writer
func (p *Plugins) YAML(w io.Writer) error {
	return displayYAML(w, p.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *Plugins) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
