package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Users is a list of the associated cloud resource with additional functionality
type Users struct {
	resource
	filterable
	items []types.User
}

// NewUsers constructs a new Users wrapping the given cloud type
func NewUsers(items []types.User) *Users {
	return &Users{
		resource: resource{
			name:     "user",
			plural:   "users",
			aliases:  []string{"usr", "usrs"},
			listView: true,
		},
		items: items,
	}
}

// User constructs a new Users with no underlying items, useful for
// interacting with the metadata itself.
func User() *Users {
	return NewUsers(nil)
}

func (c *Users) columns() []string {
	return []string{
		"ID",
		"Name",
		"Email",
		"Added At",
	}
}

// Table outputs the table representation to the given writer
func (c *Users) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, user := range c.items {
		table.Append([]string{
			string(user.ID),
			*user.Name,
			*user.Email,
			convert.UnixTimeMSToString(*user.AddedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *Users) JSON(w io.Writer) error {
	return displayJSON(w, c.items, c.resource.listView)
}

// YAML outputs the YAML representation to the given writer
func (c *Users) YAML(w io.Writer) error {
	return displayYAML(w, c.items, c.resource.listView)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *Users) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}
