package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/cloud/rest"
)

// NodesGetter is the getter for nodes
type NodesGetter interface {
	Nodes() NodeInterface
}

// NodeInterface is the interface for nodes
type NodeInterface interface {
	Create(*types.Node) (*types.Node, error)
	Get(id string) (*types.Node, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.Node, error)
}

// nodes implements NodeInterface
type nodes struct {
	client         rest.Interface
	organizationID string
	clusterID      string
	nodePoolID     string
}

func newNodes(c *Client, organizationID, clusterID, nodePoolID string) *nodes {
	return &nodes{
		client:         c.RESTClient(),
		organizationID: organizationID,
		clusterID:      clusterID,
		nodePoolID:     nodePoolID,
	}
}

// Create creates a node
func (c *nodes) Create(*types.Node) (*types.Node, error) {
	// TODO
	return nil, nil
}

// Get gets a node
func (c *nodes) Get(id string) (*types.Node, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/nodes/%s", c.organizationID, c.clusterID, c.nodePoolID, id)
	var out types.Node
	return &out, c.client.Get(path, &out)
}

// Delete deletes a node
func (c *nodes) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/nodes/%s", c.organizationID, c.clusterID, c.nodePoolID, id)
	return c.client.Delete(path)
}

// List lists all nodes
func (c *nodes) List() ([]types.Node, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/node-pools/%s/nodes", c.organizationID, c.clusterID, c.nodePoolID)
	out := make([]types.Node, 0)
	return out, c.client.Get(path, &out)
}
