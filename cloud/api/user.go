package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// UsersGetter is the getter for users
type UsersGetter interface {
	Users(organizationID string) UserInterface
}

// UserInterface is the interface for users
type UserInterface interface {
	Create(*types.User) (*types.User, error)
	Get(id string) (*types.User, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.User, error)
}

// users implements UserInterface
type users struct {
	client         rest.Interface
	organizationID string
}

func newUsers(c *Client, organizationID string) *users {
	return &users{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a user
func (c *users) Create(*types.User) (*types.User, error) {
	// TODO
	return nil, nil
}

// Get gets a user
func (c *users) Get(id string) (*types.User, error) {
	path := fmt.Sprintf("/v3/organizations/%s/users/%s", c.organizationID, id)
	var out types.User
	return &out, c.client.Get(path, &out)
}

// Delete deletes a user
func (c *users) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/users/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all users
func (c *users) List() ([]types.User, error) {
	path := fmt.Sprintf("/v3/organizations/%s/users", c.organizationID)
	out := make([]types.User, 0)
	return out, c.client.Get(path, &out)
}
