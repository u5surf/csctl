package api

import (
	"github.com/containership/csctl/cloud/api/types"
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
	// TODO make REST client
	// client rest.Interface
	client *Client
}

func newAccount(c *Client) *account {
	return &account{
		// TODO make REST client
		// client: c.RESTClient(),
		client: c,
	}
}

// Get gets an account
func (c *account) Get() (result *types.Account, err error) {
	// TODO RESTClient
	path := "/v3/account"
	var out types.Account
	return &out, c.client.GetResource(path, &out)
}
