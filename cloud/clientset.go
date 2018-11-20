package cloud

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/api"
	"github.com/containership/csctl/cloud/provision"
	"github.com/containership/csctl/cloud/rest"
)

// Interface is the main top-level interface to cloud
type Interface interface {
	API() api.Interface
	Provision() provision.Interface
}

// Clientset implements Interface
type Clientset struct {
	api       *api.Client
	provision *provision.Client
}

// Config is the configuration for a Clientset
type Config struct {
	Token string

	// Optional
	APIBaseURL       string
	ProvisionBaseURL string
}

// API returns an instance of the API client
func (c *Clientset) API() api.Interface {
	return c.api
}

// Provision returns an instance of the Provision client
func (c *Clientset) Provision() provision.Interface {
	return c.provision
}

// New constructs a new Clientset with the given config
// If base URLs are not provided, they will be defaulted by the underlying clients
func New(cfg Config) (*Clientset, error) {
	api, err := api.New(rest.Config{
		BaseURL: cfg.APIBaseURL,
		Token:   cfg.Token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "constructing API client")
	}

	provision, err := provision.New(rest.Config{
		BaseURL: cfg.ProvisionBaseURL,
		Token:   cfg.Token,
	})
	if err != nil {
		return nil, errors.Wrap(err, "constructing provision client")
	}

	return &Clientset{
		api:       api,
		provision: provision,
	}, nil
}
