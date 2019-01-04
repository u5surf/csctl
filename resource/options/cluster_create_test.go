package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClusterCreateDefaultAndValidate(t *testing.T) {
	var opts = ClusterCreate{}
	// Not everything is defaultable, so empty opts is not ok
	err := opts.DefaultAndValidate()
	assert.Error(t, err, "empty opts not ok")

	opts = ClusterCreate{
		TemplateID:  "1234",
		ProviderID:  "4321",
		Name:        "name",
		Environment: "env",
	}
	err = opts.DefaultAndValidate()
	assert.NoError(t, err, "good opts")
	// TODO make labels user-appendable and test that
	assert.NotEmpty(t, opts.labels, "labels set")

	noTemplate := opts
	noTemplate.TemplateID = ""
	err = noTemplate.DefaultAndValidate()
	assert.Error(t, err, "template is required")

	noProvider := opts
	noProvider.ProviderID = ""
	err = noProvider.DefaultAndValidate()
	assert.Error(t, err, "provider is required")

	noName := opts
	noName.Name = ""
	err = noName.DefaultAndValidate()
	assert.Error(t, err, "name is required")

	noEnvironment := opts
	noEnvironment.Environment = ""
	err = noEnvironment.DefaultAndValidate()
	assert.Error(t, err, "environment is required")
}
