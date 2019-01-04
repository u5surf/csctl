package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validUUID = "528cdd1c-871c-45cb-bef9-a839dcaa4795"
)

func TestPacketDefaultAndValidate(t *testing.T) {
	opts := PacketTemplateCreate{
		ProjectID: validUUID,
	}
	err := opts.DefaultAndValidate()
	assert.NoError(t, err, "only required options ok")

	opts.ProjectID = ""
	err = opts.validateProjectID()
	assert.Error(t, err, "missing required option not ok")

	// Fields that are not user-settable but required obtain a value
	assert.NotEmpty(t, opts.providerName, "provider name set")
}

func TestPacketTemplate(t *testing.T) {
	opts := PacketTemplateCreate{
		ProjectID: validUUID,
		Facility:  "ewr1",
		Plan:      "t1.small.x86",
	}
	err := opts.DefaultAndValidate()
	assert.NoError(t, err)

	req := opts.CreateTemplateRequest()
	assert.NoError(t, req.Validate(nil), "valid request created")
}

func TestValidateProjectID(t *testing.T) {
	opts := PacketTemplateCreate{
		ProjectID: validUUID,
	}
	err := opts.validateProjectID()
	assert.NoError(t, err)

	opts.ProjectID = ""
	err = opts.validateProjectID()
	assert.Error(t, err, "project is required")
}

func TestDefaultAndValidateFacility(t *testing.T) {
	opts := PacketTemplateCreate{}
	err := opts.defaultAndValidateFacility()
	assert.NoError(t, err)
	assert.NotEmpty(t, opts.Facility, "facility set")
}

func TestDefaultAndValidatePlan(t *testing.T) {
	opts := PacketTemplateCreate{}
	err := opts.defaultAndValidatePlan()
	assert.NoError(t, err)
	assert.NotEmpty(t, opts.Plan, "plan set")
}
