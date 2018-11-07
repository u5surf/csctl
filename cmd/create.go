package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/resource"
)

// TODO this is terrible. See other TODOs.
var opts resource.DigitalOceanTemplateCreateOptions

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
	Long: `Create a resource

TODO this is a long description`,

	Args: cobra.RangeArgs(1, 2),

	Run: func(cmd *cobra.Command, args []string) {
		resourceName := args[0]
		switch {
		case resource.Template().HasAlias(resourceName):
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			switch opts.ProviderName {
			case "digitalocean", "digital_ocean":
				if err := opts.DefaultAndValidate(); err != nil {
					fmt.Printf("Error validating options: %s\n", err)
					return
				}

				t := opts.Template()

				// TODO get response
				err := clientset.Provision().Templates(organizationID).Create(&t)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("Template created successfully!")

			case "google", "amazon_web_services", "azure", "packet":
				fmt.Printf("Error: provider %s not yet implemented\n", opts.ProviderName)
			default:
				fmt.Printf("Error: invalid provider name specified: %q\n", opts.ProviderName)
			}

		default:
			fmt.Printf("Error: invalid resource name specified: %q\n", resourceName)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// No defaulting is performed here because the logic in many cases is nontrivial,
	// and we'd like to be consistent with where and how we default.
	createCmd.Flags().StringVarP(&opts.ProviderName, "provider", "p", "", "provider name")
	createCmd.Flags().Int32VarP(&opts.MasterCount, "master-count", "m", 0, "number of nodes in master node pool")
	createCmd.Flags().Int32VarP(&opts.WorkerCount, "worker-count", "w", 0, "number of nodes in worker node pool")

	createCmd.Flags().StringVar(&opts.MasterKubernetesVersion, "master-kubernetes-version", "", "Kubernetes version for master node pool")
	createCmd.Flags().StringVar(&opts.WorkerKubernetesVersion, "worker-kubernetes-version", "", "Kubernetes version for worker node pool")

	createCmd.Flags().StringVar(&opts.Description, "description", "", "template description")

	// DigitalOcean
	// TODO this is terrible, leverage cobra flag sets or make digitalocea
	// subcommand...or something
	// TODO print default values
	createCmd.Flags().StringVar(&opts.Image, "image", "", "droplet image")
	createCmd.Flags().StringVar(&opts.Region, "region", "", "region")
	createCmd.Flags().StringVar(&opts.InstanceSize, "size", "", "instance size")
}
