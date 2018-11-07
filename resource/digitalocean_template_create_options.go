package resource

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

type DigitalOceanTemplateCreateOptions struct {
	TemplateCreateOptions

	// Defaultable
	Image        string
	Region       string
	InstanceSize string

	// Not user-settable; always defaulted
	backups           bool
	monitoring        bool
	privateNetworking bool
}

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

	return nil
}

func (o *DigitalOceanTemplateCreateOptions) Template() types.Template {
	return types.Template{
		ProviderName: &o.ProviderName,
		Description:  &o.Description,
		Engine:       &o.engine,

		Configuration: &types.TemplateConfiguration{
			Resource: &types.TemplateResource{
				DigitaloceanDroplet: types.DigitalOceanDropletMap{
					// TODO master and worker different
					o.MasterNodePoolName: o.DigitalOceanDropletConfiguration(),
					o.WorkerNodePoolName: o.DigitalOceanDropletConfiguration(),
				},
			},

			Variable: o.NodePoolVariableMap(),
		},
	}
}

func (o *DigitalOceanTemplateCreateOptions) DigitalOceanDropletConfiguration() types.DigitalOceanDropletConfiguration {
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
