package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/api/types"
)

// AccessTokenCreate is the set of options required to create an access token
type AccessTokenCreate struct {
	Name string
}

// DefaultAndValidate defaults and validates all options
func (o *AccessTokenCreate) DefaultAndValidate() error {
	if o.Name == "" {
		return errors.New("please specify name with --name")
	}

	return nil
}

// CreateAccessTokenRequest constructs an AccessToken from these options
func (o *AccessTokenCreate) CreateAccessTokenRequest() types.CreateAccessTokenRequest {
	return types.CreateAccessTokenRequest{
		Name: &o.Name,
	}
}
