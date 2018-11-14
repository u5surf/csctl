package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// getPluginCatalogCmd represents the getPluginCatalog command
var getPluginCatalogCmd = &cobra.Command{
	Use:     "plugin-catalog",
	Short:   "Get the plugin catalog",
	Aliases: resource.PluginCatalog().Aliases(),

	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := clientset.API().PluginCatalog().Get()
		if err != nil {
			return err
		}

		plugins := resource.NewPluginCatalog(v)

		outputResponse(plugins)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getPluginCatalogCmd)
}
