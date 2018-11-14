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
	path := fmt.Sprintf("/v3/plugins")
	var out types.PluginCatalog
	return &out, c.client.Get(path, &out)
}
