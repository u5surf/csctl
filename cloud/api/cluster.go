package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// ClustersGetter is the getter for clusters
type ClustersGetter interface {
	Clusters(organizationID string) ClusterInterface
}

// ClusterInterface is the interface for clusters
type ClusterInterface interface {
	Create(*types.Cluster) (*types.Cluster, error)
	Get(id string) (*types.Cluster, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.Cluster, error)
}

// clusters implements ClusterInterface
type clusters struct {
	client         rest.Interface
	organizationID string
}

func newClusters(c *Client, organizationID string) *clusters {
	return &clusters{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a cluster
func (c *clusters) Create(*types.Cluster) (*types.Cluster, error) {
	// TODO
	return nil, nil
}

// Get gets a cluster
func (c *clusters) Get(id string) (*types.Cluster, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", c.organizationID, id)
	var out types.Cluster
	return &out, c.client.Get(path, &out)
}

// Delete deletes a cluster
func (c *clusters) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all clusters
func (c *clusters) List() ([]types.Cluster, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters", c.organizationID)
	out := make([]types.Cluster, 0)
	return out, c.client.Get(path, &out)
}
