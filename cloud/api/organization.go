package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// OrganizationsGetter is the getter for organizations
type OrganizationsGetter interface {
	Organizations() OrganizationInterface
}

// OrganizationInterface is the interface for organizations
type OrganizationInterface interface {
	Create(*types.Organization) (*types.Organization, error)
	Get(id string) (*types.Organization, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.Organization, error)
}

// organizations implements OrganizationInterface
type organizations struct {
	client rest.Interface
}

func newOrganizations(c *Client) *organizations {
	return &organizations{
		client: c.RESTClient(),
	}
}

// Create creates an organization
func (c *organizations) Create(*types.Organization) (*types.Organization, error) {
	// TODO
	return nil, nil
}

// Get gets an organization
func (c *organizations) Get(id string) (*types.Organization, error) {
	path := fmt.Sprintf("/v3/organizations/%s", id)
	var out types.Organization
	return &out, c.client.Get(path, &out)
}

// Delete deletes an organization
func (c *organizations) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s", id)
	return c.client.Delete(path)
}

// List lists all organizations
func (c *organizations) List() ([]types.Organization, error) {
	path := "/v3/organizations"
	out := make([]types.Organization, 0)
	return out, c.client.Get(path, &out)
}
