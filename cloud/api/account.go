package api

import (
	"github.com/containership/csctl/cloud/api/types"
)

type AccountGetter interface {
	Account() AccountInterface
}

type AccountInterface interface {
	Get() (*types.Account, error)
}

// account implements AccountInterface
type account struct {
	// TODO make REST client
	// client rest.Interface
	client *APIClient
}

func newAccount(c *APIClient) *account {
	return &account{
		// TODO make REST client
		// client: c.RESTClient(),
		client: c,
	}
}

func (c *account) Get() (result *types.Account, err error) {
	// TODO RESTClient
	path := "/v3/account"
	var out types.Account
	return &out, c.client.GetResource(path, &out)
	return nil, nil
}
