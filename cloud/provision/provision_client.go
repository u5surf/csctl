package provision

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ProvisionInterface interface {
	CKEClustersGetter
	NodePoolsGetter
}

type ProvisionClient struct {
	name string
	// TODO make REST client and remove baseURL below
	// restClient RESTClient
	baseURL *url.URL
	token   string
}

type Config struct {
	Token   string
	BaseURL url.URL
}

// TODO don't depend on env / viper here
// Semi-weird because config comes through clientset
func New(cfg *Config) (*ProvisionClient, error) {
	baseURL, err := url.Parse(viper.GetString("provisionBaseURL"))
	if err != nil {
		return nil, errors.Wrap(err, "parsing provisionBaseURL")
	}

	return &ProvisionClient{
		name:    "Provision",
		baseURL: baseURL,
		token:   cfg.Token,
	}, nil
}

func (c *ProvisionClient) CKEClusters(organizationID string) CKEClusterInterface {
	return newCKEClusters(c, organizationID)
}

func (c *ProvisionClient) NodePools(organizationID, clusterID string) NodePoolInterface {
	return newNodePools(c, organizationID, clusterID)
}

// TODO RESTClient
func (c *ProvisionClient) GetResource(path string, output interface{}) error {
	url, _ := c.baseURL.Parse(path)

	authHeader := fmt.Sprintf("JWT %s", c.token)

	fmt.Println(url)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(output).
		Get(url.String())

	if err != nil {
		return errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership %s Service responded with status %d: %s",
			c.name, resp.StatusCode(), resp.Body())
	}

	return nil
}

func (c *ProvisionClient) DeleteResource(path string) error {
	url, _ := c.baseURL.Parse(path)

	authHeader := fmt.Sprintf("JWT %s", c.token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		Delete(url.String())

	if err != nil {
		return errors.Wrap(err, "error deleting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership Cloud responded with status %d: %s", resp.StatusCode(), resp.Body())
	}

	return nil
}
