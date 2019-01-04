package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/plugin"
)

var (
	validClusterCreateOpts = ClusterCreate{
		TemplateID:  validUUID,
		ProviderID:  validUUID,
		Name:        "name",
		Environment: "env",
	}
)

func TestDigitalOceanClusterCreateDefaultAndValidate(t *testing.T) {
	emptyOpts := DigitalOceanClusterCreate{}
	err := emptyOpts.DefaultAndValidate()
	assert.Error(t, err, "empty parent opts is not ok")

	opts := DigitalOceanClusterCreate{ClusterCreate: validClusterCreateOpts}

	err = opts.DefaultAndValidate()
	assert.NoError(t, err, "empty DO opts is ok")
}

func TestCreateCKEClusterRequest(t *testing.T) {
	emptyOpts := DigitalOceanClusterCreate{}
	req := emptyOpts.CreateCKEClusterRequest()
	assert.NotNil(t, req, "CKE cluster request is never nil")
}

func TestDigitalOceanDefaultAndValidateCNI(t *testing.T) {
	opts := DigitalOceanClusterCreate{ClusterCreate: validClusterCreateOpts}

	err := opts.defaultAndValidateCNI()
	assert.NoError(t, err, "empty CNI flag is ok")

	opts.PluginCNIFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCNI()
	assert.Error(t, err, "invalid CNI plugin flag")

	opts.PluginCNIFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCNI()
	assert.Error(t, err, "disabling CNI is not allowed")
}

func TestDigitalOceanDefaultAndValidateCCM(t *testing.T) {
	opts := DigitalOceanClusterCreate{ClusterCreate: validClusterCreateOpts}

	err := opts.defaultAndValidateCCM()
	assert.NoError(t, err, "empty CCM flag is ok")

	opts.PluginCCMFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCCM()
	assert.Error(t, err, "invalid CCM plugin flag")

	opts.PluginCCMFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCCM()
	assert.NoError(t, err, "disabling CCM is allowed")
}

func TestDigitalOceanDefaultAndValidateCSI(t *testing.T) {
	opts := DigitalOceanClusterCreate{ClusterCreate: validClusterCreateOpts}

	err := opts.defaultAndValidateCSI()
	assert.NoError(t, err, "empty CSI flag is ok")

	opts.PluginCSIFlag = plugin.Flag{Val: "=invalid"}
	err = opts.defaultAndValidateCSI()
	assert.Error(t, err, "invalid CSI plugin flag")

	opts.PluginCSIFlag = plugin.Flag{Val: "none"}
	err = opts.defaultAndValidateCSI()
	assert.NoError(t, err, "disabling CSI is allowed")
}
