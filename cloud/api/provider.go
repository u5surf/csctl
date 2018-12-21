package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// ProvidersGetter is the getter for providers
type ProvidersGetter interface {
	Providers(organizationID string) ProviderInterface
}

// ProviderInterface is the interface for providers
type ProviderInterface interface {
	Create(*types.Provider) (*types.Provider, error)
	Get(id string) (*types.Provider, error)
	Delete(id string) error
	List() ([]types.Provider, error)
}

// providers implements ProviderInterface
type providers struct {
	client         rest.Interface
	organizationID string
}

func newProviders(c *Client, organizationID string) *providers {
	return &providers{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a provider
func (c *providers) Create(*types.Provider) (*types.Provider, error) {
	// TODO
	return nil, nil
}

// Get gets a provider
func (c *providers) Get(id string) (*types.Provider, error) {
	path := fmt.Sprintf("/v3/organizations/%s/providers/%s", c.organizationID, id)
	var out types.Provider
	return &out, c.client.Get(path, &out)
}

// Delete deletes a provider
func (c *providers) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/providers/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all providers
func (c *providers) List() ([]types.Provider, error) {
	path := fmt.Sprintf("/v3/organizations/%s/providers", c.organizationID)
	out := make([]types.Provider, 0)
	return out, c.client.Get(path, &out)
}
