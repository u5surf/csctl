package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var packetCreateTemplateOpts options.PacketTemplateCreate

// createTemplatePacketCmd represents the createTemplateDigitalOcean command
var createTemplatePacketCmd = &cobra.Command{
	Use:     "packet",
	Short:   "Create a Packet template",
	Args:    cobra.NoArgs,
	Aliases: []string{"pkt"},

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		packetCreateTemplateOpts.TemplateCreate = createTemplateOpts

		if err := packetCreateTemplateOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := packetCreateTemplateOpts.CreateTemplateRequest()

		template, err := clientset.Provision().Templates(organizationID).Create(&req)
		if err != nil {
			return err
		}

		t := resource.NewTemplates([]types.Template{*template})
		return t.Table(os.Stdout)
	},
}

func init() {
	createTemplateCmd.AddCommand(createTemplatePacketCmd)

	bindCommandToOrganizationScope(createTemplatePacketCmd, false)

	createTemplatePacketCmd.Flags().StringVar(&packetCreateTemplateOpts.ProjectID, "project", "", "Packet project ID")
	createTemplatePacketCmd.Flags().StringVar(&packetCreateTemplateOpts.Facility, "facility", "", "facility")
	createTemplatePacketCmd.Flags().StringVar(&packetCreateTemplateOpts.Plan, "plan", "", "plan")
}
