package auth

import (
	"fmt"

	"github.com/containership/csctl/cloud/auth/types"
	"github.com/containership/csctl/cloud/rest"
)

// LoginGetter is the getter for login
type LoginGetter interface {
	Login(method string) LoginInterface
}

// LoginInterface is the interface for login
type LoginInterface interface {
	Post(*types.LoginRequest) (*types.AccountToken, error)
}

// login implements LoginInterface
type login struct {
	client rest.Interface
	method string
}

func newLogin(c *Client, method string) *login {
	return &login{
		client: c.RESTClient(),
		method: method,
	}
}

// Post attempts to log in the user and generate an account token
func (c *login) Post(req *types.LoginRequest) (*types.AccountToken, error) {
	path := fmt.Sprintf("/v3/login/%s", c.method)
	var out types.AccountToken
	return &out, c.client.Post(path, req, &out)
}
