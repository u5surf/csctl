package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/cloud/rest"
)

// NodePoolsGetter is the getter for node pools
type NodePoolsGetter interface {
	NodePools(organizationID, clusterID string) NodePoolInterface
}

// NodePoolInterface is the interface for node pools
type NodePoolInterface interface {
	Create(*types.NodePool) (*types.NodePool, error)
	Get(id string) (*types.NodePool, error)
	Delete(id string) error
	List() ([]types.NodePool, error)
	Scale(id string, req *types.ScaleNodePoolRequest) (*types.NodePool, error)
}

// nodePools implements NodePoolInterface
type nodePools struct {
	client         rest.Interface
	organizationID string
	clusterID      string
}

func newNodePools(c *Client, organizationID, clusterID string) *nodePools {
	return &nodePools{
		client:         c.RESTClient(),
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
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s", c.organizationID, c.clusterID, id)
	var out types.NodePool
	return &out, c.client.Get(path, &out)
}

// Delete deletes a node pool
func (c *nodePools) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s", c.organizationID, c.clusterID, id)
	return c.client.Delete(path)
}

// List lists all node pools
func (c *nodePools) List() ([]types.NodePool, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools", c.organizationID, c.clusterID)
	out := make([]types.NodePool, 0)
	return out, c.client.Get(path, &out)
}

// Scale scales a node pool
func (c *nodePools) Scale(id string, req *types.ScaleNodePoolRequest) (*types.NodePool, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s", c.organizationID, c.clusterID, id)
	var out types.NodePool
	return &out, c.client.Patch(path, req, &out)
}
