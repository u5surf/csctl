package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

type organizations struct {
	resource
	filterable
	items []types.Organization
}

func NewOrganizations(items []types.Organization) *organizations {
	return &organizations{
		resource: resource{
			name:    "organization",
			plural:  "organizations",
			aliases: []string{"org", "orgs"},
		},
		items: items,
	}
}

func Organizations() *organizations {
	return NewOrganizations(nil)
}

func (o *organizations) columns() []string {
	return []string{
		"Name",
		"ID",
		"Owner ID",
		"Created At",
	}
}

func (o *organizations) Table(w io.Writer) error {
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

func (o *organizations) JSON(w io.Writer) error {
	return displayJSON(w, o.items)
}

func (o *organizations) YAML(w io.Writer) error {
	return displayYAML(w, o.items)
}

func (o *organizations) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.items)
}

func (o *organizations) FilterByOwnerID(id string) {
	filtered := make([]types.Organization, 0)
	for _, cluster := range o.items {
		if string(cluster.OwnerID) == id {
			filtered = append(filtered, cluster)
		}
	}

	o.items = filtered
}
