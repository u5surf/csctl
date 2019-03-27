package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Accounts is a list of the associated cloud resource with additional functionality
type Accounts struct {
	resource
	items []types.Account
}

// NewAccounts constructs a new Accounts wrapping the given cloud type
func NewAccounts(items []types.Account) *Accounts {
	return &Accounts{
		resource: resource{
			name:     "account",
			aliases:  []string{"acct"},
			listView: true,
		},
		items: items,
	}
}

// Account constructs a new Accounts with no underlying items, useful for
// interacting with the metadata itself.
func Account() *Accounts {
	return NewAccounts(nil)
}

func (a *Accounts) columns() []string {
	return []string{
		"Name",
		"ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (a *Accounts) Table(w io.Writer) error {
	table := table.New(w, a.columns())

	for _, acct := range a.items {
		table.Append([]string{
			*acct.Name,
			string(acct.ID),
			convert.UnixTimeMSToString(*acct.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (a *Accounts) JSON(w io.Writer) error {
	return displayJSON(w, a.items, a.resource.listView)
}

// YAML outputs the YAML representation to the given writer
func (a *Accounts) YAML(w io.Writer) error {
	return displayYAML(w, a.items, a.resource.listView)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (a *Accounts) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, a.items)
}
