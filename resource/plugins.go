package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

type plugins struct {
	resource
	filterable
	items []types.Plugin
}

func NewPlugins(items []types.Plugin) *plugins {
	return &plugins{
		resource: resource{
			name:    "plugin",
			plural:  "plugins",
			aliases: []string{"plug", "plugs", "plgn", "plgns"},
		},
		items: items,
	}
}

func Plugins() *plugins {
	return NewPlugins(nil)
}

func (p *plugins) columns() []string {
	return []string{
		"ID",
		"Type",
		"Implementation",
		"Version",
		"Created At",
	}
}

func (p *plugins) Table(w io.Writer) error {
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

func (p *plugins) JSON(w io.Writer) error {
	return displayJSON(w, p.items)
}

func (p *plugins) YAML(w io.Writer) error {
	return displayYAML(w, p.items)
}

func (p *plugins) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
