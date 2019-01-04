package options

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/resource/plugin"
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

func TestDefaultAndValidateMetrics(t *testing.T) {
	opts := ClusterCreate{}

	err := opts.defaultAndValidateMetrics()
	assert.NoError(t, err, "empty metrics flag is ok")

	opts.PluginMetricsFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateMetrics()
	assert.Error(t, err, "invalid metrics plugin flag")

	opts.PluginMetricsFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateMetrics()
	assert.NoError(t, err, "disabling metrics is allowed")
}

func TestDefaultAndValidateClusterManagement(t *testing.T) {
	opts := ClusterCreate{}

	err := opts.defaultAndValidateClusterManagement()
	assert.NoError(t, err, "empty cluster management flag is ok")

	opts.PluginClusterManagementFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateClusterManagement()
	assert.Error(t, err, "invalid cluster management plugin flag")

	opts.PluginClusterManagementFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateClusterManagement()
	assert.Error(t, err, "disabling cluster management is not allowed")

	opts.PluginClusterManagementFlag = plugin.Flag{Val: "notcontainership"}
	err = opts.defaultAndValidateClusterManagement()
	assert.Error(t, err, "only containership cluster management plugin is allowed")
}

func TestDefaultAndValidateAutoscaler(t *testing.T) {
	opts := ClusterCreate{}

	err := opts.defaultAndValidateAutoscaler()
	assert.NoError(t, err, "empty metrics flag is ok")

	opts.PluginAutoscalerFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateAutoscaler()
	assert.Error(t, err, "invalid metrics plugin flag")

	opts.PluginAutoscalerFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateAutoscaler()
	assert.NoError(t, err, "disabling metrics is allowed")
}
