package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
)

// NodePoolsGetter is the getter for node pools
type NodePoolsGetter interface {
	NodePools() NodePoolInterface
}

// NodePoolInterface is the interface for node pools
type NodePoolInterface interface {
	Create(*types.NodePool) (*types.NodePool, error)
	Get(id string) (*types.NodePool, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.NodePool, error)
}

// nodePools implements NodePoolInterface
type nodePools struct {
	// TODO make REST client
	// client rest.Interface
	client         *Client
	organizationID string
	clusterID      string
}

func newNodePools(c *Client, organizationID, clusterID string) *nodePools {
	return &nodePools{
		// TODO make REST client
		// client: c.RESTClient(),
		client:         c,
		organizationID: organizationID,
		clusterID:      clusterID,
	}
}

// Create creates a node pool
func (c *nodePools) Create(*types.NodePool) (*types.NodePool, error) {
	// TODO
	return nil, nil
}

// Get gets a node pool
func (c *nodePools) Get(id string) (*types.NodePool, error) {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s", c.organizationID, c.clusterID, id)
	var out types.NodePool
	return &out, c.client.GetResource(path, &out)
}

// Delete deletes a node pool
func (c *nodePools) Delete(id string) error {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s", c.organizationID, c.clusterID, id)
	return c.client.DeleteResource(path)
}

// List lists all node pools
func (c *nodePools) List() ([]types.NodePool, error) {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools", c.organizationID, c.clusterID)
	out := make([]types.NodePool, 0)
	return out, c.client.GetResource(path, &out)
}
