package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
)

type OrganizationsGetter interface {
	Organizations() OrganizationInterface
}

type OrganizationInterface interface {
	Create(*types.Organization) (*types.Organization, error)
	Get(id string) (*types.Organization, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.Organization, error)
}

// organizations implements OrganizationInterface
type organizations struct {
	// TODO make REST client
	// client rest.Interface
	client *APIClient
}

func newOrganizations(c *APIClient) *organizations {
	return &organizations{
		// TODO make REST client
		// client: c.RESTClient(),
		client: c,
	}
}

func (c *organizations) Create(*types.Organization) (*types.Organization, error) {
	// TODO
	return nil, nil
}

func (c *organizations) Get(id string) (*types.Organization, error) {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s", id)
	var out types.Organization
	return &out, c.client.GetResource(path, &out)
}

func (c *organizations) Delete(id string) error {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s", id)
	return c.client.DeleteResource(path)
}

func (c *organizations) List() ([]types.Organization, error) {
	// TODO RESTClient
	path := "/v3/organizations"
	out := make([]types.Organization, 0)
	return out, c.client.GetResource(path, &out)
}
