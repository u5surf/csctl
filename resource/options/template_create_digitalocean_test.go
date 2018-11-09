package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDODefaultAndValidate(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{}
	// Everything is defaultable, so no error should occur for empty opts
	err := opts.DefaultAndValidate()
	assert.Nil(t, err, "empty opts ok")

	// Fields that are not user-settable but required obtain a value
	assert.NotEmpty(t, opts.providerName, "provider name set")
}

func TestDOTemplate(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{
		Image:        "ubuntu-16-04-x64",
		Region:       "nyc2",
		InstanceSize: "s-2vcpu-2gb",
	}
	err := opts.DefaultAndValidate()
	assert.Nil(t, err)

	tmpl := opts.Template()
	assert.Nil(t, tmpl.Validate(nil), "valid template created")
}

func TestDefaultAndValidateImage(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{}
	err := opts.defaultAndValidateImage()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.Image, "image set")
}

func TestDefaultAndValidateRegion(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{}
	err := opts.defaultAndValidateRegion()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.Region, "region set")
}

func TestDefaultAndValidateInstanceSize(t *testing.T) {
	var opts = DigitalOceanTemplateCreate{}
	err := opts.defaultAndValidateInstanceSize()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.InstanceSize, "instance size set")
}
