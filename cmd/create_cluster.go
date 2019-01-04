package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
	"github.com/containership/csctl/resource/plugin"
)

var createClusterOpts options.ClusterCreate

// createClusterCmd represents the createCluster command
var createClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create (provision) a CKE cluster",
}

func init() {
	createCmd.AddCommand(createClusterCmd)

	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.TemplateID, "template", "t", "", "template ID to create from")
	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.ProviderID, "provider", "p", "", "provider ID (credentials) to use for provisioning")
	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.Name, "name", "n", "", "cluster name")
	createClusterCmd.PersistentFlags().StringVarP(&createClusterOpts.Environment, "environment", "e", "", "environment")

	// Plugins which are never provider-specific
	createClusterDigitalOceanCmd.Flags().StringVar(&createClusterOpts.PluginMetricsFlag.Val, "plugin-metrics", "",
		fmt.Sprintf("metrics plugin (specify %q to disable)", plugin.NoImplementation))
	createClusterDigitalOceanCmd.Flags().StringVar(&createClusterOpts.PluginClusterManagementFlag.Val, "plugin-cluster-management", "",
		"Cluster Management plugin (implementation must be \"containership\")")
	createClusterDigitalOceanCmd.Flags().StringVar(&createClusterOpts.PluginAutoscalerFlag.Val, "plugin-autoscaler", "",
		"autoscaler plugin")
}
