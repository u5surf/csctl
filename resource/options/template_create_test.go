package options

import (
	"testing"

	"github.com/Masterminds/semver"

	"github.com/stretchr/testify/assert"
)

func TestTemplateCreateDefaultAndValidate(t *testing.T) {
	var opts = TemplateCreate{}
	// Everything is defaultable, so no error should occur for empty opts
	err := opts.DefaultAndValidate()
	assert.Nil(t, err, "empty opts ok")

	// Fields that are not user-settable but required obtain a value
	assert.NotEmpty(t, opts.engine, "engine set")
	assert.NotEmpty(t, opts.masterMode, "master mode set")
	assert.NotEmpty(t, opts.workerMode, "worker mode set")
	assert.NotEmpty(t, opts.nodePoolType, "node pool type set")
}

func TestNodePoolVariableMap(t *testing.T) {
	var opts = TemplateCreate{
		MasterCount:             3,
		WorkerCount:             5,
		MasterKubernetesVersion: "1.11.1",
		Description:             "testing",
		MasterNodePoolName:      "master",
	}
	err := opts.DefaultAndValidate()
	assert.Nil(t, err)

	m := opts.NodePoolVariableMap()
	assert.Nil(t, m.Validate(nil), "valid node pool map created")
	// TODO support multiple node pools, don't hardcode
	assert.Len(t, m, 2, "correct number of node pools")
}

func TestDefaultAndValidateMasterCount(t *testing.T) {
	var opts = TemplateCreate{}
	err := opts.defaultAndValidateMasterCount()
	assert.Nil(t, err)
	assert.True(t, opts.MasterCount > 0 && opts.MasterCount != 2, "default master count valid")

	opts.MasterCount = 2
	err = opts.defaultAndValidateMasterCount()
	assert.NotNil(t, err, "invalid master count")

	opts.MasterCount = -1
	err = opts.defaultAndValidateMasterCount()
	assert.NotNil(t, err, "negative master count")
}

func TestDefaultAndValidateWorkerCount(t *testing.T) {
	var opts = TemplateCreate{}
	err := opts.defaultAndValidateWorkerCount()
	assert.Nil(t, err)
	assert.True(t, opts.WorkerCount > 0, "default worker count valid")

	opts.WorkerCount = -1
	err = opts.defaultAndValidateWorkerCount()
	assert.NotNil(t, err, "negative worker count")
}

func TestDefaultAndValidateMasterNodePoolName(t *testing.T) {
	var opts = TemplateCreate{}
	err := opts.defaultAndValidateMasterNodePoolName()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.MasterNodePoolName, "default to non-empty string")
}

func TestDefaultAndValidateWorkerNodePoolName(t *testing.T) {
	var opts = TemplateCreate{}
	err := opts.defaultAndValidateWorkerNodePoolName()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.WorkerNodePoolName, "default to non-empty string")
}

func TestDefaultAndValidateKubernetesVersions(t *testing.T) {
	var opts = TemplateCreate{}
	err := opts.defaultAndValidateKubernetesVersions()
	assert.Nil(t, err)
	_, err = semver.NewVersion(opts.MasterKubernetesVersion)
	assert.Nil(t, err, "default master valid semver")

	_, err = semver.NewVersion(opts.WorkerKubernetesVersion)
	assert.Nil(t, err, "default worker valid semver")

	opts = TemplateCreate{MasterKubernetesVersion: "1.10.1"}
	err = opts.defaultAndValidateKubernetesVersions()
	assert.Nil(t, err)
	assert.Equal(t, opts.MasterKubernetesVersion, opts.WorkerKubernetesVersion,
		"default worker version to master version")
}

func TestDefaultAndValidateDescription(t *testing.T) {
	var opts = TemplateCreate{}
	err := opts.defaultAndValidateDescription()
	assert.Nil(t, err)
	assert.NotEmpty(t, opts.Description, "default to non-empty string")
}
