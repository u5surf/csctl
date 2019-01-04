package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource/options"
	//"github.com/containership/csctl/resource/plugin"
)

var doCreateClusterOpts options.DigitalOceanClusterCreate

// createClusterDigitalOceanCmd represents the createClusterDigitalOcean command
var createClusterDigitalOceanCmd = &cobra.Command{
	Use:     "digitalocean",
	Short:   "Create a DigitalOcean cluster",
	Args:    cobra.NoArgs,
	Aliases: []string{"do"},

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		doCreateClusterOpts.ClusterCreate = createClusterOpts

		if err := doCreateClusterOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := doCreateClusterOpts.CreateCKEClusterRequest()

		resp, err := clientset.Provision().CKEClusters(organizationID).Create(&req)
		if err != nil {
			return err
		}

		fmt.Printf("Cluster %s provisioning initiated successfully\n", resp.ID)
		return nil
	},
}

func init() {
	createClusterCmd.AddCommand(createClusterDigitalOceanCmd)

	bindCommandToOrganizationScope(createClusterDigitalOceanCmd, false)
}
