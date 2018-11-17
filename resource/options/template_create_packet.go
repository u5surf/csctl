package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

// PacketTemplateCreate is the set of options required
// to create a Packet template
type PacketTemplateCreate struct {
	TemplateCreate

	// Required
	ProjectID string

	// Defaultable
	Facility string
	Plan     string

	// Not user-settable
	providerName string
}

// DefaultAndValidate defaults and validates all options
func (o *PacketTemplateCreate) DefaultAndValidate() error {
	if err := o.TemplateCreate.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}

	if err := o.validateProjectID(); err != nil {
		return errors.Wrap(err, "validating Packet project ID")
	}

	if err := o.defaultAndValidateFacility(); err != nil {
		return errors.Wrap(err, "validating facility")
	}

	if err := o.defaultAndValidatePlan(); err != nil {
		return errors.Wrap(err, "validating instance size")
	}

	o.providerName = "packet"

	return nil
}

// CreateTemplateRequest constructs a CreateTemplateRequest from these options
func (o *PacketTemplateCreate) CreateTemplateRequest() types.CreateTemplateRequest {
	return types.CreateTemplateRequest{
		ProviderName: &o.providerName,
		Description:  &o.Description,
		Engine:       &o.engine,

		Configuration: &types.TemplateConfiguration{
			Resource: &types.TemplateResource{
				PacketDevice: types.PacketDeviceMap{
					// TODO master and worker different
					o.MasterNodePoolName: o.digitalOceanDropletConfiguration(),
					o.WorkerNodePoolName: o.digitalOceanDropletConfiguration(),
				},
			},

			Variable: o.NodePoolVariableMap(),
		},
	}
}

func (o *PacketTemplateCreate) digitalOceanDropletConfiguration() types.PacketDeviceConfiguration {
	return types.PacketDeviceConfiguration{
		ProjectID: types.UUID(o.ProjectID),
		Facility:  &o.Facility,
		Plan:      &o.Plan,
	}
}

func (o *PacketTemplateCreate) validateProjectID() error {
	if o.ProjectID == "" {
		return errors.New("Packet project ID must be provided")
	}

	return nil
}

func (o *PacketTemplateCreate) defaultAndValidateFacility() error {
	if o.Facility == "" {
		o.Facility = "ewr1"
	}

	return nil
}

func (o *PacketTemplateCreate) defaultAndValidatePlan() error {
	if o.Plan == "" {
		o.Plan = "t1.small.x86"
	}

	return nil
}
