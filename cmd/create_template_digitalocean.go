package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
	"github.com/containership/csctl/resource/options"
)

var doCreateTemplateOpts options.DigitalOceanTemplateCreate

// createTemplateDigitalOceanCmd represents the createTemplateDigitalOcean command
var createTemplateDigitalOceanCmd = &cobra.Command{
	Use:     "digitalocean",
	Short:   "Create a DigitalOcean template",
	Args:    cobra.NoArgs,
	Aliases: []string{"do"},

	PreRunE: orgScopedPreRunE,

	RunE: func(cmd *cobra.Command, args []string) error {
		doCreateTemplateOpts.TemplateCreate = createTemplateOpts

		if err := doCreateTemplateOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		req := doCreateTemplateOpts.CreateTemplateRequest()

		template, err := clientset.Provision().Templates(organizationID).Create(&req)
		if err != nil {
			return err
		}

		t := resource.NewTemplates([]types.Template{*template})
		return t.Table(os.Stdout)
	},
}

func init() {
	createTemplateCmd.AddCommand(createTemplateDigitalOceanCmd)

	bindCommandToOrganizationScope(createTemplateDigitalOceanCmd, false)

	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.Image, "image", "", "droplet image")
	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.Region, "region", "", "region")
	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.InstanceSize, "size", "", "instance size")
}
