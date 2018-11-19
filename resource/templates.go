package resource

import (
	"io"

	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/pkg/convert"
	"github.com/containership/csctl/resource/table"
)

// Templates is a list of the associated cloud resource with additional functionality
type Templates struct {
	resource
	filterable
	items []types.Template
}

// filterFunc returns true if the item should be filtered in
// or false if it should be excluded (removed from the slice)
// TODO this doesn't belong here once generic filtering is implemented
// See filter.go
type filterFunc func(types.Template) bool

// NewTemplates constructs a new Templates wrapping the given cloud type
func NewTemplates(items []types.Template) *Templates {
	return &Templates{
		resource: resource{
			name:    "template",
			plural:  "templates",
			aliases: []string{"tmpl", "tmpls"},
		},
		items: items,
	}
}

// Template constructs a new Templates with no underlying items, useful for
// interacting with the metadata itself.
func Template() *Templates {
	return NewTemplates(nil)
}

func (c *Templates) columns() []string {
	return []string{
		"ID",
		"Description",
		"Provider Name",
		"Master Version",
		"Owner ID",
		"Created At",
	}
}

// Table outputs the table representation to the given writer
func (c *Templates) Table(w io.Writer) error {
	table := table.New(w, c.columns())

	for _, tmpl := range c.items {
		masterVersion, err := getMasterKubernetesVersion(&tmpl)
		if err != nil {
			return errors.Wrapf(err, "retrieving master version for template %q", string(tmpl.ID))
		}

		var desc = "none"
		if tmpl.Description != nil && *tmpl.Description != "" {
			desc = *tmpl.Description
		}

		table.Append([]string{
			string(tmpl.ID),
			desc,
			*tmpl.ProviderName,
			masterVersion,
			string(tmpl.OwnerID),
			convert.UnixTimeMSToString(*tmpl.CreatedAt),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (c *Templates) JSON(w io.Writer) error {
	return displayJSON(w, c.items)
}

// YAML outputs the YAML representation to the given writer
func (c *Templates) YAML(w io.Writer) error {
	return displayYAML(w, c.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (c *Templates) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, c.items)
}

// FilterByOwnerID filters the underlying items by owner ID
func (c *Templates) FilterByOwnerID(id string) {
	c.applyFilter(func(t types.Template) bool {
		return string(t.OwnerID) == id
	})
}

// FilterByEngine filters the underlying items by engine
func (c *Templates) FilterByEngine(engine string) {
	c.applyFilter(func(t types.Template) bool {
		return t.Engine != nil && *t.Engine == engine
	})
}

func (c *Templates) applyFilter(f filterFunc) {
	filtered := make([]types.Template, 0)
	for _, tmpl := range c.items {
		if f(tmpl) {
			filtered = append(filtered, tmpl)
		}
	}

	c.items = filtered
}

// getMasterKubernetesVersion returns the Kubernetes version of the master node pool
// for the given template, or an error
// TODO the Template type is nasty to interact with due to how the API response
// is structured. We should provide convenience functions in the official go client itself.
func getMasterKubernetesVersion(t *types.Template) (string, error) {
	if t == nil {
		return "", errors.New("template is nil")
	}
	if t.Configuration == nil || t.Configuration.Variable == nil {
		return "", errors.New("template configuration is nil")
	}

	for _, np := range t.Configuration.Variable {
		if np.Default.KubernetesMode != nil &&
			*np.Default.KubernetesMode == types.NodePoolKubernetesModeMaster {
			return *np.Default.KubernetesVersion, nil
		}
	}

	return "", errors.New("could not find master node pool")
}
