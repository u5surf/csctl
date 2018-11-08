package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

var doCreateTemplateOpts resource.DigitalOceanTemplateCreateOptions

// createTemplateDigitalOceanCmd represents the createTemplateDigitalOcean command
var createTemplateDigitalOceanCmd = &cobra.Command{
	Use:   "digitalocean",
	Short: "Create a DigitalOcean template",
	Long: `TODO

TODO this is a long description`,
	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO how to avoid setting parent field?
		doCreateTemplateOpts.TemplateCreateOptions = createTemplateOpts

		if err := doCreateTemplateOpts.DefaultAndValidate(); err != nil {

			fmt.Printf("Error validating options: %s\n", err)
			return
		}

		t := doCreateTemplateOpts.Template()

		// TODO get response
		err := clientset.Provision().Templates(organizationID).Create(&t)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Template created successfully!")
	},
}

func init() {
	createTemplateCmd.AddCommand(createTemplateDigitalOceanCmd)

	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.Image, "image", "", "droplet image")
	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.Region, "region", "", "region")
	createTemplateDigitalOceanCmd.Flags().StringVar(&doCreateTemplateOpts.InstanceSize, "size", "", "instance size")
}
