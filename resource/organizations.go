package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Organizations is a list of the associated cloud resource with additional functionality
type Organizations struct {
	resource
	filterable
	items []types.Organization
}

// NewOrganizations constructs a new Organizations wrapping the given cloud type
func NewOrganizations(items []types.Organization) *Organizations {
	return &Organizations{
		resource: resource{
			name:    "organization",
			plural:  "Organizations",
			aliases: []string{"org", "orgs"},
		},
		items: items,
	}
}

// Organization constructs a new Organizations with no underlying items, useful for
// interacting with the metadata itself.
func Organization() *Organizations {
	return NewOrganizations(nil)
}

func (o *Organizations) columns() []string {
	return []string{
		"Name",
		"ID",
		"Owner ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (o *Organizations) Table(w io.Writer) error {
	table := table.New(w, o.columns())

	for _, org := range o.items {
		table.Append([]string{
			*org.Name,
			string(org.ID),
			string(org.OwnerID),
			convert.UnixTimeMSToString(*org.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (o *Organizations) JSON(w io.Writer) error {
	return displayJSON(w, o.items)
}

// YAML outputs the YAML representation to the given writer
func (o *Organizations) YAML(w io.Writer) error {
	return displayYAML(w, o.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (o *Organizations) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.items)
}

// FilterByOwnerID filters the underlying items by owner ID
func (o *Organizations) FilterByOwnerID(id string) {
	filtered := make([]types.Organization, 0)
	for _, cluster := range o.items {
		if string(cluster.OwnerID) == id {
			filtered = append(filtered, cluster)
		}
	}

	o.items = filtered
}
