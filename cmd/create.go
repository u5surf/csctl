package cmd

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	provisiontypes "github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource"
)

// Flags
var (
	// Required
	providerName string

	// Defaultable
	masterCount int32
	workerCount int32

	masterKubernetesVersion string
	workerKubernetesVersion string

	image        string
	region       string
	instanceSize string

	templateDescription string
)

type templateCreateOptions struct {
	// Inherited and required
	organizationID string

	// Required
	providerName string

	// Defaultable
	masterCount int32
	workerCount int32

	masterKubernetesVersion string
	workerKubernetesVersion string

	masterSchedulable bool

	description string
}

type templateDigitalOceanOptions struct {
	templateCreateOptions

	// Defaultable
	image        string
	region       string
	instanceSize string
}

func (o *templateCreateOptions) defaultAndValidate() error {
	if err := o.validateProviderName(); err != nil {
		return err
	}

	if err := o.defaultAndValidateMasterCount(); err != nil {
		return errors.Wrap(err, "master count")
	}

	if err := o.defaultAndValidateWorkerCount(); err != nil {
		return errors.Wrap(err, "worker count")
	}

	if err := o.defaultAndValidateKubernetesVersions(); err != nil {
		return errors.Wrap(err, "worker count")
	}

	return nil
}

func (o *templateCreateOptions) defaultAndValidateMasterCount() error {
	if o.masterCount == 0 {
		o.masterCount = 1
		return nil
	}
	if o.masterCount < 1 || o.masterCount == 2 {
		return errors.New("master count must be 1 or >= 3")
	}
	return nil
}

func (o *templateCreateOptions) defaultAndValidateWorkerCount() error {
	if o.workerCount == 0 {
		o.workerCount = 1
		return nil
	}
	if o.workerCount < 1 {
		return errors.New("worker count must be >= 1")
	}
	return nil
}

func (o *templateCreateOptions) defaultAndValidateKubernetesVersions() error {
	if o.masterKubernetesVersion == "" {
		o.masterKubernetesVersion = "1.12.1"
	}
	mv, err := semver.NewVersion(o.masterKubernetesVersion)
	if err != nil {
		return errors.Wrap(err, "master semver")
	}
	// Note that String() returns the version with the leading 'v' stripped
	// if applicable, which is what we want for cloud interactions.
	o.masterKubernetesVersion = mv.String()

	if o.workerKubernetesVersion == "" {
		// Worker pools default to master version, which is validated by now
		o.workerKubernetesVersion = o.masterKubernetesVersion
		return nil
	}
	mw, err := semver.NewVersion(o.workerKubernetesVersion)
	if err != nil {
		return errors.Wrap(err, "worker semver")
	}
	o.workerKubernetesVersion = mw.String()

	return nil
}

func (o *templateCreateOptions) validateProviderName() error {
	switch o.providerName {
	case "digital_ocean", "google", "amazon_web_services", "azure", "packet":
		break
	case "":
		return errors.Errorf("provider name is required")
	}
	return nil
}

func (o *templateDigitalOceanOptions) defaultAndValidate() error {
	if err := o.defaultAndValidateImage(); err != nil {
		return errors.Wrap(err, "validating image name")
	}

	if err := o.defaultAndValidateRegion(); err != nil {
		return errors.Wrap(err, "validating region")
	}

	if err := o.defaultAndValidateInstanceSize(); err != nil {
		return errors.Wrap(err, "validating instance size")
	}

	return nil
}

func (o *templateDigitalOceanOptions) defaultAndValidateImage() error {
	// TODO client-side validation, maybe
	if o.image == "" {
		o.image = "ubuntu-16-04-x64"
	}

	return nil
}

func (o *templateDigitalOceanOptions) defaultAndValidateRegion() error {
	// TODO client-side validation, maybe
	if o.region == "" {
		o.region = "nyc1"
	}

	return nil
}

func (o *templateDigitalOceanOptions) defaultAndValidateInstanceSize() error {
	// TODO client-side validation, maybe
	if o.instanceSize == "" {
		o.instanceSize = "s-1vcpu-2gb"
	}

	return nil
}

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

			createOpts := templateCreateOptions{
				providerName: providerName,

				masterCount: masterCount,
				workerCount: workerCount,

				masterKubernetesVersion: masterKubernetesVersion,
				workerKubernetesVersion: workerKubernetesVersion,

				description: templateDescription,
			}

			if err := createOpts.defaultAndValidate(); err != nil {
				fmt.Printf("Error validating options: %s\n", err)
				return
			}

			switch createOpts.providerName {
			case "digital_ocean":
				opts := templateDigitalOceanOptions{
					templateCreateOptions: createOpts,

					image:        image,
					region:       region,
					instanceSize: instanceSize,
				}

				if err := opts.defaultAndValidate(); err != nil {
					fmt.Printf("Error validating DigitalOcean options: %s\n", err)
					return
				}

				engine := "containership_kubernetes_engine"
				masterMode := "master"
				t := provisiontypes.Template{
					ProviderName: &opts.providerName,
					Description:  &opts.description,
					Engine:       &engine,

					Configuration: &provisiontypes.TemplateConfiguration{
						Variable: provisiontypes.TemplateVariableMap{
							"master-pool": provisiontypes.TemplateVariableDefault{
								Default: &provisiontypes.TemplateNodePool{
									Count:             &opts.masterCount,
									Etcd:              true,
									IsSchedulable:     true,
									KubernetesMode:    &masterMode,
									KubernetesVersion: &opts.masterKubernetesVersion,
								},
							},
						},
					},
				}

				err := clientset.Provision().Templates(organizationID).Create(&t)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("Template created successfully!")

			case "google", "amazon_web_services", "azure", "packet":
				fmt.Printf("Error: provider %s not yet implemented\n", providerName)
			default:
				fmt.Printf("Error: invalid provider name specified: %q\n", providerName)
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
	createCmd.Flags().StringVarP(&providerName, "provider", "p", "", "provider name")
	createCmd.Flags().Int32VarP(&masterCount, "master-count", "m", 0, "number of nodes in master node pool")
	createCmd.Flags().Int32VarP(&workerCount, "worker-count", "w", 0, "number of nodes in worker node pool")

	createCmd.Flags().StringVar(&masterKubernetesVersion, "master-kubernetes-version", "", "Kubernetes version for master node pool")
	createCmd.Flags().StringVar(&workerKubernetesVersion, "worker-kubernetes-version", "", "Kubernetes version for worker node pool")

	createCmd.Flags().StringVar(&templateDescription, "description", "None", "template description")
}
