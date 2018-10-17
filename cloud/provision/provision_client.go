package provision

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/rest"
)

// Interface is the interface for Provision
type Interface interface {
	RESTClient() rest.Interface
	CKEClustersGetter
	NodePoolsGetter
}

// Client is the Provision client
type Client struct {
	name       string
	restClient *rest.Client
}

// New constructs a new Provision client
func New(cfg *rest.Config) (*Client, error) {
	restClient, err := rest.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "constructing REST client")
	}

	return &Client{
		name:       "Provision",
		restClient: restClient,
	}, nil
}

// RESTClient returns the REST client associated with this client
func (c *Client) RESTClient() rest.Interface {
	return c.restClient
}

// CKEClusters returns the CKE clusters interface
func (c *Client) CKEClusters(organizationID string) CKEClusterInterface {
	return newCKEClusters(c, organizationID)
}

// NodePools returns the node pools interface
func (c *Client) NodePools(organizationID, clusterID string) NodePoolInterface {
	return newNodePools(c, organizationID, clusterID)
}
