package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Providers is a list of the associated cloud resource with additional functionality
type Providers struct {
	resource
	items []types.Provider
}

// NewProviders constructs a new Providers wrapping the given cloud type
func NewProviders(items []types.Provider) *Providers {
	return &Providers{
		resource: resource{
			name:   "provider",
			plural: "providers",
		},
		items: items,
	}
}

// Provider constructs a new Providers with no underlying items, useful for
// interacting with the metadata itself.
func Provider() *Providers {
	return NewProviders(nil)
}

func (c *Providers) columns() []string {
	return []string{
		"ID",
		"Description",
		"Provider Name",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (c *Providers) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, provider := range c.items {
		table.Append([]string{
			string(provider.ID),
			*provider.Description,
			*provider.Provider,
			convert.UnixTimeMSToString(*provider.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *Providers) JSON(w io.Writer) error {
	return displayJSON(w, c.items)
}

// YAML outputs the YAML representation to the given writer
func (c *Providers) YAML(w io.Writer) error {
	return displayYAML(w, c.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *Providers) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}
