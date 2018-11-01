package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

type accounts struct {
	resource
	items []types.Account
}

func NewAccounts(items []types.Account) *accounts {
	return &accounts{
		resource: resource{
			name:    "account",
			aliases: []string{"acct"},
		},
		items: items,
	}
}

func Accounts() *accounts {
	return NewAccounts(nil)
}

func (a *accounts) columns() []string {
	return []string{
		"Name",
		"ID",
		"Created At",
	}
}

func (a *accounts) Table(w io.Writer) error {
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

func (a *accounts) JSON(w io.Writer) error {
	return displayJSON(w, a.items)
}

func (a *accounts) YAML(w io.Writer) error {
	return displayYAML(w, a.items)
}

func (a *accounts) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, a.items)
}
