package auth

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/rest"
)

const (
	defaultBaseURL = "https://auth.containership.io"
)

// Interface is the interface for Auth
type Interface interface {
	RESTClient() rest.Interface
	LoginGetter
}

// Client is the Auth client
type Client struct {
	name       string
	restClient *rest.Client
}

// New constructs a new API client
func New(cfg rest.Config) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}

	restClient, err := rest.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "constructing REST client")
	}

	return &Client{
		name:       "Auth",
		restClient: restClient,
	}, nil
}

// RESTClient returns the REST client associated with this client
func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

// Login returns the login interface
func (c *Client) Login(method string) LoginInterface {
	return newLogin(c, method)
}
