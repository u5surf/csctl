package api

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/rest"
)

// Interface is the interface for API
type Interface interface {
	RESTClient() rest.Interface
	OrganizationsGetter
	AccountGetter
}

// Client is the API client
type Client struct {
	name       string
	restClient *rest.Client
}

// New constructs a new API client
func New(cfg *rest.Config) (*Client, error) {
	restClient, err := rest.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "constructing REST client")
	}

	return &Client{
		name:       "API",
		restClient: restClient,
	}, nil
}

// RESTClient returns the REST client associated with this client
func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

// Organizations returns the organizations interface
func (c *Client) Organizations() OrganizationInterface {
	return newOrganizations(c)
}

// Account returns the account interface
func (c *Client) Account() AccountInterface {
	return newAccount(c)
}
