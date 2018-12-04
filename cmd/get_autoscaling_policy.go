package cmd

import (
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

// getAutoscalingPolicyCmd represents the getAutoscalingPolicy command
var getAutoscalingPolicyCmd = &cobra.Command{
	Use:     "autoscaling-policy",
	Short:   "Get an autoscaling policy or list of autoscaling policies",
	Aliases: resource.AutoscalingPolicy().Aliases(),

	Args: cobra.MaximumNArgs(1),

	PreRunE: clusterScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		var resp = make([]types.AutoscalingPolicy, 1)
		var err error
		if len(args) == 1 {
			var v *types.AutoscalingPolicy
			v, err = clientset.Provision().AutoscalingPolicies(organizationID, clusterID).Get(args[0])
			resp[0] = *v
		} else {
			resp, err = clientset.Provision().AutoscalingPolicies(organizationID, clusterID).List()
		}

		if err != nil {
			return err
		}

		nps := resource.NewAutoscalingPolicies(resp)
		outputResponse(nps)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getAutoscalingPolicyCmd)

	bindCommandToClusterScope(getAutoscalingPolicyCmd, false)
}
