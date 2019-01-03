package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// RegistryGetter is the getter for registries
type RegistryGetter interface {
	Registries(organizationID string) RegistryInterface
}

// RegistryInterface is the interface for registries
type RegistryInterface interface {
	Create(*types.Registry) (*types.Registry, error)
	Get(id string) (*types.Registry, error)
	Delete(id string) error
	List() ([]types.Registry, error)
}

// registry implements RegistryInterface
type registry struct {
	client         rest.Interface
	organizationID string
}

func newRegistry(c *Client, organizationID string) *registry {
	return &registry{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a registry
func (c *registry) Create(*types.Registry) (*types.Registry, error) {
	// TODO
	return nil, nil
}

// Get gets a registry
func (c *registry) Get(id string) (*types.Registry, error) {
	path := fmt.Sprintf("/v3/organizations/%s/registries/%s", c.organizationID, id)
	var out types.Registry
	return &out, c.client.Get(path, &out)
}

// Delete deletes a registry
func (c *registry) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/registries/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all registries
func (c *registry) List() ([]types.Registry, error) {
	path := fmt.Sprintf("/v3/organizations/%s/registries", c.organizationID)
	out := make([]types.Registry, 0)
	return out, c.client.Get(path, &out)
}
