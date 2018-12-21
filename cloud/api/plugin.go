package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// PluginsGetter is the getter for plugins
type PluginsGetter interface {
	Plugins(organizationID, clusterID string) PluginInterface
}

// PluginInterface is the interface for plugins
type PluginInterface interface {
	Create(*types.Plugin) (*types.Plugin, error)
	Get(id string) (*types.Plugin, error)
	Delete(id string) error
	List() ([]types.Plugin, error)
}

// plugins implements PluginInterface
type plugins struct {
	client         rest.Interface
	organizationID string
	clusterID      string
}

func newPlugins(c *Client, organizationID, clusterID string) *plugins {
	return &plugins{
		client:         c.RESTClient(),
		organizationID: organizationID,
		clusterID:      clusterID,
	}
}

// Create creates a plugin
func (c *plugins) Create(*types.Plugin) (*types.Plugin, error) {
	// TODO
	return nil, nil
}

// Get gets a plugin
func (c *plugins) Get(id string) (*types.Plugin, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/plugins/%s", c.organizationID, c.clusterID, id)
	var out types.Plugin
	return &out, c.client.Get(path, &out)
}

// Delete deletes a plugin
func (c *plugins) Delete(id string) error {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/plugins/%s", c.organizationID, c.clusterID, id)
	return c.client.Delete(path)
}

// List lists all plugins
func (c *plugins) List() ([]types.Plugin, error) {
	path := fmt.Sprintf("/v3/organizations/%s/clusters/%s/plugins", c.organizationID, c.clusterID)
	out := make([]types.Plugin, 0)
	return out, c.client.Get(path, &out)
}
