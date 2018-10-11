package cloud

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/api"
	"github.com/containership/csctl/cloud/provision"
)

// Clientset is a set of clients for interacting with Containership Cloud
type Clientset struct {
	api       *api.Client
	provision *provision.Client
	//proxy     *proxy.ProxyClient
}

// Config is the configuration for a Clientset
type Config struct {
	Token string
}

// API returns an instance of the API client
func (c *Clientset) API() *api.Client {
	return c.api
}

// Provision returns an instance of the Provision client
func (c *Clientset) Provision() *provision.Client {
	return c.provision
}

// New constructs a new Clientset
func New(cfg *Config) (*Clientset, error) {
	api, err := api.New(&api.Config{
		Token: cfg.Token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "constructing API client")
	}

	provision, err := provision.New(&provision.Config{
		Token: cfg.Token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "constructing provision client")
	}

	return &Clientset{
		api:       api,
		provision: provision,
	}, nil
}
