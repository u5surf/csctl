package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

var doCreateTemplateOpts resource.DigitalOceanTemplateCreateOptions

// createTemplateDigitalOceanCmd represents the createTemplateDigitalOcean command
var createTemplateDigitalOceanCmd = &cobra.Command{
	Use:   "digitalocean",
	Short: "Create a DigitalOcean template",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		doCreateTemplateOpts.TemplateCreateOptions = createTemplateOpts

		if err := doCreateTemplateOpts.DefaultAndValidate(); err != nil {
			return errors.Wrap(err, "validating options")
		}

		t := doCreateTemplateOpts.Template()

		newTemplate, err := clientset.Provision().Templates(organizationID).Create(&t)
		if err != nil {
			return err
		}

		templates := resource.NewTemplates([]types.Template{*newTemplate})
		return templates.Table(os.Stdout)
	},
}

func init() {
	createTemplateCmd.AddCommand(createTemplateDigitalOceanCmd)

	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.Image, "image", "", "droplet image")
	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.Region, "region", "", "region")
	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.InstanceSize, "size", "", "instance size")
}
