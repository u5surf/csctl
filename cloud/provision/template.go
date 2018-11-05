package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/cloud/rest"
)

// TemplatesGetter is the getter for templates
type TemplatesGetter interface {
	Templates() TemplateInterface
}

// TemplateInterface is the interface for templates
type TemplateInterface interface {
	Create(*types.Template) (*types.Template, error)
	Get(id string) (*types.Template, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.Template, error)
}

// templates implements TemplateInterface
type templates struct {
	client         rest.Interface
	organizationID string
}

func newTemplates(c *Client, organizationID string) *templates {
	return &templates{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a template
func (t *templates) Create(*types.Template) (*types.Template, error) {
	// TODO
	return nil, nil
}

// Get gets a template
func (t *templates) Get(id string) (*types.Template, error) {
	path := fmt.Sprintf("/v3/organizations/%s/templates/%s", t.organizationID, id)
	var out types.Template
	return &out, t.client.Get(path, &out)
}

// Delete deletes a template
func (t *templates) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/templates/%s", t.organizationID, id)
	return t.client.Delete(path)
}

// List lists all templates
func (t *templates) List() ([]types.Template, error) {
	path := fmt.Sprintf("/v3/organizations/%s/templates", t.organizationID)
	out := make([]types.Template, 0)
	return out, t.client.Get(path, &out)
}
