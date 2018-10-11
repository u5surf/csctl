package provision

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Interface is the interface for Provision
type Interface interface {
	CKEClustersGetter
	NodePoolsGetter
}

// Client is the Provision client
type Client struct {
	name string
	// TODO make REST client and remove baseURL below
	// restClient RESTClient
	baseURL *url.URL
	token   string
}

// Config is the configuration for Provision client
type Config struct {
	Token   string
	BaseURL url.URL
}

// New constructs a new Provision client
// TODO don't depend on env / viper here
// Semi-weird because config comes through clientset
func New(cfg *Config) (*Client, error) {
	baseURL, err := url.Parse(viper.GetString("provisionBaseURL"))
	if err != nil {
		return nil, errors.Wrap(err, "parsing provisionBaseURL")
	}

	return &Client{
		name:    "Provision",
		baseURL: baseURL,
		token:   cfg.Token,
	}, nil
}

// CKEClusters returns the CKE clusters interface
func (c *Client) CKEClusters(organizationID string) CKEClusterInterface {
	return newCKEClusters(c, organizationID)
}

// NodePools returns the node pools interface
func (c *Client) NodePools(organizationID, clusterID string) NodePoolInterface {
	return newNodePools(c, organizationID, clusterID)
}

// GetResource gets a resource at the given path and stores the result in output
// or returns an error
// TODO RESTClient
func (c *Client) GetResource(path string, output interface{}) error {
	url, _ := c.baseURL.Parse(path)

	authHeader := fmt.Sprintf("JWT %s", c.token)

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

// DeleteResource deletes the resource at the given path or returns an error
func (c *Client) DeleteResource(path string) error {
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
