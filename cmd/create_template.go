package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/containership/csctl/resource"
)

var createTemplateOpts resource.TemplateCreateOptions

// createTemplateCmd represents the createTemplate command
var createTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a cluster template for provisioning",
	Args:  cobra.NoArgs,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		organizationID = viper.GetString("organization")
		if organizationID == "" {
			return errors.New("please specify an organization via --organization or config file")
		}

		return nil
	},
}

func init() {
	createCmd.AddCommand(createTemplateCmd)

	bindCommandToOrganizationScope(createTemplateCmd, true)

	// No defaulting is performed here because the logic in many cases is nontrivial,
	// and we'd like to be consistent with where and how we default.
	createTemplateCmd.Flags().Int32VarP(&createTemplateOpts.MasterCount, "master-count", "m", 0, "number of nodes in master node pool")
	createTemplateCmd.Flags().Int32VarP(&createTemplateOpts.WorkerCount, "worker-count", "w", 0, "number of nodes in worker node pool")

	createTemplateCmd.Flags().StringVar(&createTemplateOpts.MasterKubernetesVersion, "master-kubernetes-version", "", "Kubernetes version for master node pool")
	createTemplateCmd.Flags().StringVar(&createTemplateOpts.WorkerKubernetesVersion, "worker-kubernetes-version", "", "Kubernetes version for worker node pool")

	createTemplateCmd.Flags().StringVar(&createTemplateOpts.Description, "description", "", "template description")
}
