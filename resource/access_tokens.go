package resource

import (
	"io"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// AccessTokens is a list of the associated cloud resource with additional functionality
type AccessTokens struct {
	resource
	filterable
	items []types.AccessToken
}

// NewAccessTokens constructs a new AccessTokens wrapping the given cloud type
func NewAccessTokens(items []types.AccessToken) *AccessTokens {
	return &AccessTokens{
		resource: resource{
			name:    "access-token",
			plural:  "access-tokens",
			aliases: []string{"token", "tokens"},
		},
		items: items,
	}
}

// AccessToken constructs a new AccessTokens with no underlying items, useful for
// interacting with the metadata itself.
func AccessToken() *AccessTokens {
	return NewAccessTokens(nil)
}

func (t *AccessTokens) columns() []string {
	return []string{
		"Name",
		"ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (t *AccessTokens) Table(w io.Writer) error {
	table := table.New(w, t.columns())

	for _, token := range t.items {
		table.Append([]string{
			*token.Name,
			string(token.ID),
			convert.UnixTimeMSToString(*token.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (t *AccessTokens) JSON(w io.Writer) error {
	return displayJSON(w, t.items)
}

// YAML outputs the YAML representation to the given writer
func (t *AccessTokens) YAML(w io.Writer) error {
	return displayYAML(w, t.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (t *AccessTokens) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, t.items)
}
