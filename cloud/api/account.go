package api

import (
	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// AccountGetter is the getter for account
type AccountGetter interface {
	Account() AccountInterface
}

// AccountInterface is the interface for account
type AccountInterface interface {
	Get() (*types.Account, error)
}

// account implements AccountInterface
type account struct {
	client rest.Interface
}

func newAccount(c *Client) *account {
	return &account{
		client: c.RESTClient(),
	}
}

// Get gets an account
func (c *account) Get() (result *types.Account, err error) {
	path := "/v3/account"
	var out types.Account
	return &out, c.client.Get(path, &out)
}
