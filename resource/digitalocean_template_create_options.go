package resource

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

// DigitalOceanTemplateCreateOptions is the set of options required
// to create a DigitalOcean template
type DigitalOceanTemplateCreateOptions struct {
	TemplateCreateOptions

	// Defaultable
	Image        string
	Region       string
	InstanceSize string

	// Not user-settable; always defaulted
	providerName string

	backups           bool
	monitoring        bool
	privateNetworking bool
}

// DefaultAndValidate defaults and validates all options
func (o *DigitalOceanTemplateCreateOptions) DefaultAndValidate() error {
	if err := o.TemplateCreateOptions.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}
	if err := o.defaultAndValidateImage(); err != nil {
		return errors.Wrap(err, "validating image name")
	}

	if err := o.defaultAndValidateRegion(); err != nil {
		return errors.Wrap(err, "validating region")
	}

	if err := o.defaultAndValidateInstanceSize(); err != nil {
		return errors.Wrap(err, "validating instance size")
	}

	o.backups = false
	o.monitoring = false
	o.privateNetworking = true

	o.providerName = "digital_ocean"

	return nil
}

// Template constructs a full template from these options
func (o *DigitalOceanTemplateCreateOptions) Template() types.Template {
	return types.Template{
		ProviderName: &o.providerName,
		Description:  &o.Description,
		Engine:       &o.engine,

		Configuration: &types.TemplateConfiguration{
			Resource: &types.TemplateResource{
				DigitaloceanDroplet: types.DigitalOceanDropletMap{
					// TODO master and worker different
					o.MasterNodePoolName: o.digitalOceanDropletConfiguration(),
					o.WorkerNodePoolName: o.digitalOceanDropletConfiguration(),
				},
			},

			Variable: o.NodePoolVariableMap(),
		},
	}
}

func (o *DigitalOceanTemplateCreateOptions) digitalOceanDropletConfiguration() types.DigitalOceanDropletConfiguration {
	return types.DigitalOceanDropletConfiguration{
		Image:  &o.Image,
		Region: &o.Region,
		Size:   &o.InstanceSize,

		Backups:           o.backups,
		Monitoring:        o.monitoring,
		PrivateNetworking: &o.privateNetworking,
	}
}

func (o *DigitalOceanTemplateCreateOptions) defaultAndValidateImage() error {
	if o.Image == "" {
		o.Image = "ubuntu-16-04-x64"
	}

	return nil
}

func (o *DigitalOceanTemplateCreateOptions) defaultAndValidateRegion() error {
	if o.Region == "" {
		o.Region = "nyc1"
	}

	return nil
}

func (o *DigitalOceanTemplateCreateOptions) defaultAndValidateInstanceSize() error {
	if o.InstanceSize == "" {
		o.InstanceSize = "s-1vcpu-2gb"
	}

	return nil
}
