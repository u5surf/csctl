package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/plugin"
)

// DigitalOceanClusterCreate is the set of options required
// to create a DigitalOcean cluster
type DigitalOceanClusterCreate struct {
	ClusterCreate

	PluginCNIFlag               plugin.Flag
	PluginCSIFlag               plugin.Flag
	PluginCCMFlag               plugin.Flag
	PluginMetricsFlag           plugin.Flag
	PluginClusterManagementFlag plugin.Flag
	PluginAutoscalerFlag        plugin.Flag

	plugins []*types.CreateCKEClusterPlugin
}

// DefaultAndValidate defaults and validates all options
func (o *DigitalOceanClusterCreate) DefaultAndValidate() error {
	if err := o.ClusterCreate.DefaultAndValidate(); err != nil {
		return errors.Wrap(err, "validating generic create options")
	}

	if err := o.defaultAndValidateCNI(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCNI)
	}

	if err := o.defaultAndValidateCSI(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCSI)
	}

	if err := o.defaultAndValidateCCM(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeCCM)
	}

	if err := o.defaultAndValidateMetrics(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeMetrics)
	}

	if err := o.defaultAndValidateClusterManagement(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeClusterManagement)
	}

	if err := o.defaultAndValidateAutoscaler(); err != nil {
		return errors.Wrapf(err, "validating %s plugin", plugin.TypeAutoscaler)
	}

	return nil
}

// CreateCKEClusterRequest constructs a CreateCKEClusterRequest from these options
func (o *DigitalOceanClusterCreate) CreateCKEClusterRequest() types.CreateCKEClusterRequest {
	return types.CreateCKEClusterRequest{
		ProviderID: types.UUID(o.ProviderID),
		TemplateID: types.UUID(o.TemplateID),
		Labels:     o.labels,
		Plugins:    o.plugins,
	}
}

func (o *DigitalOceanClusterCreate) defaultAndValidateCNI() error {
	impl, version, err := o.PluginCNIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return errors.Errorf("CNI plugin is required (can't specify an implementation of %q)", plugin.NoImplementation)
	}

	if impl != "" && impl != "calico" {
		return errors.New("only calico CNI implementation is allowed")
	}
	impl = "calico"

	pType := "cni"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *DigitalOceanClusterCreate) defaultAndValidateCCM() error {
	impl, version, err := o.PluginCCMFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != "digitalocean" {
		return errors.New("only digitalocean CCM implementation is allowed")
	}
	impl = "digitalocean"

	pType := "cloud_controller_manager"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
		Configuration: map[string]string{
			"provider_id": o.ProviderID,
		},
	})

	return nil
}

func (o *DigitalOceanClusterCreate) defaultAndValidateCSI() error {
	impl, version, err := o.PluginCSIFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != "digitalocean" {
		return errors.New("only digitalocean CCM implementation is allowed")
	}
	impl = "digitalocean"

	pType := "csi"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
		Configuration: map[string]string{
			"provider_id": o.ProviderID,
		},
	})

	return nil
}

func (o *DigitalOceanClusterCreate) defaultAndValidateMetrics() error {
	impl, version, err := o.PluginMetricsFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != "prometheus" {
		return errors.New("only prometheus metrics implementation is allowed")
	}
	impl = "prometheus"

	pType := "metrics"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *DigitalOceanClusterCreate) defaultAndValidateClusterManagement() error {
	impl, version, err := o.PluginClusterManagementFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return errors.Errorf("cluster management plugin is required (can't specify an implementation of %q)", plugin.NoImplementation)
	}

	if impl != "" && impl != "containership" {
		return errors.New("only \"containership\" cluster management implementation is allowed")
	}
	impl = "containership"

	pType := "cluster_management"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}

func (o *DigitalOceanClusterCreate) defaultAndValidateAutoscaler() error {
	impl, version, err := o.PluginAutoscalerFlag.Parse()
	if err != nil {
		return errors.Wrap(err, "parsing plugin flag")
	}

	if impl == plugin.NoImplementation {
		return nil
	}

	if impl != "" && impl != "cerebral" {
		return errors.New("only cerebral metrics implementation is allowed")
	}
	impl = "cerebral"

	pType := "autoscaler"
	o.plugins = append(o.plugins, &types.CreateCKEClusterPlugin{
		Type:           &pType,
		Implementation: &impl,
		Version:        version,
	})

	return nil
}
