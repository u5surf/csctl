package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
)

var createTemplateOpts options.TemplateCreate

// createTemplateCmd represents the createTemplate command
var createTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a cluster template for provisioning",

	// TODO using PersistentPreRunE here to org-scope all create template
	// subcommands would be nice, but it won't work currently because the
	// root command is already using it and this cobra issue is open:
	// https://github.com/spf13/cobra/issues/252
}

func init() {
	createCmd.AddCommand(createTemplateCmd)

	// No defaulting is performed here because the logic in many cases is nontrivial,
	// and we'd like to be consistent with where and how we default.
	createTemplateCmd.Flags().Int32VarP(&createTemplateOpts.MasterCount, "master-count", "m", 0, "number of nodes in master node pool")
	createTemplateCmd.Flags().Int32VarP(&createTemplateOpts.WorkerCount, "worker-count", "w", 0, "number of nodes in worker node pool")

	createTemplateCmd.Flags().StringVar(&createTemplateOpts.MasterKubernetesVersion, "master-kubernetes-version", "", "Kubernetes version for master node pool")
	createTemplateCmd.Flags().StringVar(&createTemplateOpts.WorkerKubernetesVersion, "worker-kubernetes-version", "", "Kubernetes version for worker node pool")

	createTemplateCmd.Flags().StringVar(&createTemplateOpts.Description, "description", "", "template description")
}
