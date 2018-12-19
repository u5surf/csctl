package api

import (
	"fmt"

	"github.com/containership/csctl/cloud/api/types"
	"github.com/containership/csctl/cloud/rest"
)

// PluginCatalogGetter is the getter for plugins catalog
type PluginCatalogGetter interface {
	PluginCatalog() PluginCatalogInterface
}

// PluginCatalogInterface is the interface for plugins catalog
type PluginCatalogInterface interface {
	Get() (*types.PluginCatalog, error)
	Type(t string) ([]*types.PluginDefinition, error)
	TypeImplementation(t, impl string) (*types.PluginDefinition, error)
	TypeImplementationVersion(t, impl, version string) (*types.PluginVersion, error)
}

// pluginCatalog implements PluginInterface
type pluginCatalog struct {
	client rest.Interface
}

func newPluginCatalog(c *Client) *pluginCatalog {
	return &pluginCatalog{
		client: c.RESTClient(),
	}
}

// Get plugin catalog
func (c *pluginCatalog) Get() (*types.PluginCatalog, error) {
	path := "/v3/plugins"
	var out types.PluginCatalog
	return &out, c.client.Get(path, &out)
}

// Type gets the plugin catalog filtered by type
func (c *pluginCatalog) Type(t string) ([]*types.PluginDefinition, error) {
	path := fmt.Sprintf("/v3/plugins/%s", t)
	var out []*types.PluginDefinition
	return out, c.client.Get(path, &out)
}

// TypeImplementation gets the plugin catalog filtered by type and implementation
func (c *pluginCatalog) TypeImplementation(t, impl string) (*types.PluginDefinition, error) {
	path := fmt.Sprintf("/v3/plugins/%s/implementations/%s", t, impl)
	var out types.PluginDefinition
	return &out, c.client.Get(path, &out)
}

// TypeImplementationVersion gets the plugin catalog filtered by type, implementation, and version
func (c *pluginCatalog) TypeImplementationVersion(t, impl, version string) (*types.PluginVersion, error) {
	path := fmt.Sprintf("/v3/plugins/%s/implementations/%s/versions/%s", t, impl, version)
	var out types.PluginVersion
	return &out, c.client.Get(path, &out)
}
