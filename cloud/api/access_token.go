package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// AccessTokensGetter is the getter for accessTokens
type AccessTokensGetter interface {
	AccessTokens() AccessTokenInterface
}

// AccessTokenInterface is the interface for accessTokens
type AccessTokenInterface interface {
	Create(*types.CreateAccessTokenRequest) (*types.AccessToken, error)
	Get(id string) (*types.AccessToken, error)
	Delete(id string) error
	List() ([]types.AccessToken, error)
}

// accessTokens implements AccessTokenInterface
type accessTokens struct {
	client rest.Interface
}

func newAccessTokens(t *Client) *accessTokens {
	return &accessTokens{
		client: t.RESTClient(),
	}
}

// Create creates an access token
func (t *accessTokens) Create(token *types.CreateAccessTokenRequest) (*types.AccessToken, error) {
	path := fmt.Sprintf("/v3/account/access_tokens")
	var out types.AccessToken
	return &out, t.client.Post(path, token, &out)
}

// Get gets an access token
func (t *accessTokens) Get(id string) (*types.AccessToken, error) {
	path := fmt.Sprintf("/v3/account/access_tokens/%s", id)
	var out types.AccessToken
	return &out, t.client.Get(path, &out)
}

// Delete deletes an access token
func (t *accessTokens) Delete(id string) error {
	path := fmt.Sprintf("/v3/account/access_tokens/%s", id)
	return t.client.Delete(path)
}

// List lists all access tokens
func (t *accessTokens) List() ([]types.AccessToken, error) {
	path := fmt.Sprintf("/v3/account/access_tokens")
	out := make([]types.AccessToken, 0)
	return out, t.client.Get(path, &out)
}
