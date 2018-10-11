package cloud

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/api"
	"github.com/containership/csctl/cloud/provision"
)

type Clientset struct {
	api       *api.APIClient
	provision *provision.ProvisionClient
	//proxy     *proxy.ProxyClient
}

type Config struct {
	Token string
}

func (c *Clientset) API() *api.APIClient {
	return c.api
}

func (c *Clientset) Provision() *provision.ProvisionClient {
	return c.provision
}

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
