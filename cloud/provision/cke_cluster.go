package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/cloud/rest"
)

// CKEClustersGetter is the getter for CKE clusters
type CKEClustersGetter interface {
	CKEClusters(organizationID string) CKEClusterInterface
}

// CKEClusterInterface is the interface for CKE clusters
type CKEClusterInterface interface {
	Create(*types.CKECluster) (*types.CKECluster, error)
	Get(id string) (*types.CKECluster, error)
	Delete(id string) error
	List() ([]types.CKECluster, error)
}

// ckeClusters implements CKEClusterInterface
type ckeClusters struct {
	client         rest.Interface
	organizationID string
}

func newCKEClusters(c *Client, organizationID string) *ckeClusters {
	return &ckeClusters{
		client:         c.RESTClient(),
		organizationID: organizationID,
	}
}

// Create creates a CKE cluster
func (c *ckeClusters) Create(*types.CKECluster) (*types.CKECluster, error) {
	// TODO
	return nil, nil
}

// Get gets a CKE cluster
func (c *ckeClusters) Get(id string) (*types.CKECluster, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", c.organizationID, id)
	var out types.CKECluster
	return &out, c.client.Get(path, &out)
}

// Delete deletes a CKE cluster
func (c *ckeClusters) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", c.organizationID, id)
	return c.client.Delete(path)
}

// List lists all CKE clusters
func (c *ckeClusters) List() ([]types.CKECluster, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters", c.organizationID)
	out := make([]types.CKECluster, 0)
	return out, c.client.Get(path, &out)
}
