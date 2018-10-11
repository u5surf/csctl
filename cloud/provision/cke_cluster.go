package provision

import (
	"fmt"

	"github.com/containership/csctl/cloud/provision/types"
)

type CKEClustersGetter interface {
	CKEClusters() CKEClusterInterface
}

type CKEClusterInterface interface {
	Create(*types.CKECluster) (*types.CKECluster, error)
	Get(id string) (*types.CKECluster, error)
	Delete(id string) error
	// TODO list options implemented client-side
	List() ([]types.CKECluster, error)
}

// ckeClusters implements CKEClusterInterface
type ckeClusters struct {
	// TODO make REST client
	// client rest.Interface
	client         *ProvisionClient
	organizationID string
}

func newCKEClusters(c *ProvisionClient, organizationID string) *ckeClusters {
	return &ckeClusters{
		// TODO make REST client
		// client: c.RESTClient(),
		client:         c,
		organizationID: organizationID,
	}
}

func (c *ckeClusters) Create(*types.CKECluster) (*types.CKECluster, error) {
	// TODO
	return nil, nil
}

func (c *ckeClusters) Get(id string) (*types.CKECluster, error) {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", c.organizationID, id)
	var out types.CKECluster
	return &out, c.client.GetResource(path, &out)
	return nil, nil
}

func (c *ckeClusters) Delete(id string) error {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s", c.organizationID, id)
	return c.client.DeleteResource(path)
}

func (c *ckeClusters) List() ([]types.CKECluster, error) {
	// TODO RESTClient
	path := fmt.Sprintf("/v3/organizations/%s/clusters", c.organizationID)
	out := make([]types.CKECluster, 0)
	return out, c.client.GetResource(path, &out)
}
