package provision

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/rest"
)

const (
	defaultBaseURL = "https://provision.containership.io"
)

// Interface is the interface for Provision
type Interface interface {
	RESTClient() rest.Interface
	CKEClustersGetter
	TemplatesGetter
	NodePoolsGetter
	AutoscalingPoliciesGetter
	NodesGetter
}

// Client is the Provision client
type Client struct {
	name       string
	restClient *rest.Client
}

// New constructs a new Provision client
func New(cfg rest.Config) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}

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

// Templates returns the templates interface
func (c *Client) Templates(organizationID string) TemplateInterface {
	return newTemplates(c, organizationID)
}

// NodePools returns the node pools interface
func (c *Client) NodePools(organizationID, clusterID string) NodePoolInterface {
	return newNodePools(c, organizationID, clusterID)
}

// AutoscalingPolicies returns the autoscaling policies interface
func (c *Client) AutoscalingPolicies(organizationID, clusterID string) AutoscalingPolicyInterface {
	return newAutoscalingPolicies(c, organizationID, clusterID)
}

// Nodes returns the nodes interface
func (c *Client) Nodes(organizationID, clusterID, nodePoolID string) NodeInterface {
	return newNodes(c, organizationID, clusterID, nodePoolID)
}
