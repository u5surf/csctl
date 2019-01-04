package options

import (
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/plugin"
)

// ClusterCreate is the set of options required to create a cluster
type ClusterCreate struct {
	TemplateID string
	ProviderID string

	Name        string
	Environment string

	// Plugin flags without provider-specific validation
	PluginMetricsFlag           plugin.Flag
	PluginClusterManagementFlag plugin.Flag
	PluginAutoscalerFlag        plugin.Flag

	// Flags with provider-specific validation
	PluginCNIFlag plugin.Flag
	PluginCSIFlag plugin.Flag
	PluginCCMFlag plugin.Flag

	plugins []*types.CreateCKEClusterPlugin

	// Not user-settable; always defaulted
	// TODO make user-appendable: --labels=key1=val1,key2=val2
	labels map[string]string
}

// DefaultAndValidate defaults and validates all options
func (o *ClusterCreate) DefaultAndValidate() error {
	if o.TemplateID == "" {
		return errors.New("please specify template with --template")
	}

	if o.ProviderID == "" {
		return errors.New("please specify provider credentials with --provider")
	}

	if o.Name == "" {
		return errors.New("please specify name with --name")
	}

	if o.Environment == "" {
		return errors.New("please specify name with --environment")
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

	o.labels = map[string]string{
		"containership.io/cluster-name":        o.Name,
		"containership.io/cluster-environment": o.Environment,
	}

	return nil
}

func (o *ClusterCreate) defaultAndValidateMetrics() error {
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

func (o *ClusterCreate) defaultAndValidateClusterManagement() error {
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

func (o *ClusterCreate) defaultAndValidateAutoscaler() error {
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
