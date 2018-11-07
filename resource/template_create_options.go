package resource

import (
	"github.com/Masterminds/semver"
	"github.com/pkg/errors"

	"github.com/containership/csctl/cloud/provision/types"
)

type TemplateCreateOptions struct {
	// Required
	ProviderName string

	// Defaultable
	MasterCount int32
	WorkerCount int32

	MasterKubernetesVersion string
	WorkerKubernetesVersion string

	Description string

	// TODO the following should be user-settable
	MasterNodePoolName string
	WorkerNodePoolName string

	// Not user-settable; always defaulted
	engine string

	masterMode string
	workerMode string

	nodePoolType string

	masterSchedulable bool
}

/*
type templateCreateOptionsInterface interface {
	DefaultAndValidate() error
	NodePoolVariableMap() types.TemplateVariableMap
}
*/

func (o *TemplateCreateOptions) DefaultAndValidate() error {
	if err := o.validateProviderName(); err != nil {
		return err
	}

	if err := o.defaultAndValidateMasterCount(); err != nil {
		return errors.Wrap(err, "master count")
	}

	if err := o.defaultAndValidateWorkerCount(); err != nil {
		return errors.Wrap(err, "worker count")
	}

	if err := o.defaultAndValidateMasterNodePoolName(); err != nil {
		return errors.Wrap(err, "master node pool name")
	}

	if err := o.defaultAndValidateWorkerNodePoolName(); err != nil {
		return errors.Wrap(err, "worker node pool name")
	}

	if err := o.defaultAndValidateKubernetesVersions(); err != nil {
		return errors.Wrap(err, "kubernetes versions")
	}

	if err := o.defaultAndValidateDescription(); err != nil {
		return errors.Wrap(err, "description")
	}

	o.engine = "containership_kubernetes_engine"

	o.masterMode = "master"
	o.workerMode = "worker"

	o.nodePoolType = "node_pool"

	// TODO should be user-settable
	o.masterSchedulable = true

	return nil
}

func (o *TemplateCreateOptions) NodePoolVariableMap() types.TemplateVariableMap {
	return types.TemplateVariableMap{
		o.MasterNodePoolName: types.TemplateVariableDefault{
			Default: &types.TemplateNodePool{
				Count:             &o.MasterCount,
				KubernetesMode:    &o.masterMode,
				KubernetesVersion: &o.MasterKubernetesVersion,
				Name:              &o.MasterNodePoolName,
				Type:              &o.nodePoolType,

				Etcd:          true,
				IsSchedulable: o.masterSchedulable,
			},
		},
		o.WorkerNodePoolName: types.TemplateVariableDefault{
			Default: &types.TemplateNodePool{
				Count:             &o.WorkerCount,
				KubernetesMode:    &o.workerMode,
				KubernetesVersion: &o.WorkerKubernetesVersion,
				Name:              &o.WorkerNodePoolName,
				Type:              &o.nodePoolType,
			},
		},
	}
}

func (o *TemplateCreateOptions) validateProviderName() error {
	switch o.ProviderName {
	case "digitalocean":
		// This is valid for user input since DO is technically one word,
		// but cloud expects an underscore
		o.ProviderName = "digital_ocean"
	case "digital_ocean", "google", "amazon_web_services", "azure", "packet":
		break
	case "":
		return errors.Errorf("provider name is required")
	}
	return nil
}

func (o *TemplateCreateOptions) defaultAndValidateMasterCount() error {
	if o.MasterCount == 0 {
		o.MasterCount = 1
		return nil
	}
	if o.MasterCount < 1 || o.MasterCount == 2 {
		return errors.New("master count must be 1 or >= 3")
	}
	return nil
}

func (o *TemplateCreateOptions) defaultAndValidateWorkerCount() error {
	if o.WorkerCount == 0 {
		o.WorkerCount = 1
		return nil
	}
	if o.WorkerCount < 1 {
		return errors.New("worker count must be >= 1")
	}
	return nil
}

func (o *TemplateCreateOptions) defaultAndValidateMasterNodePoolName() error {
	if o.MasterNodePoolName == "" {
		o.MasterNodePoolName = "master-pool-0"
	}
	return nil
}

func (o *TemplateCreateOptions) defaultAndValidateWorkerNodePoolName() error {
	if o.WorkerNodePoolName == "" {
		o.WorkerNodePoolName = "worker-pool-0"
	}
	return nil
}

func (o *TemplateCreateOptions) defaultAndValidateKubernetesVersions() error {
	if o.MasterKubernetesVersion == "" {
		o.MasterKubernetesVersion = "1.12.1"
	}
	mv, err := semver.NewVersion(o.MasterKubernetesVersion)
	if err != nil {
		return errors.Wrap(err, "master semver")
	}
	// Note that String() returns the version with the leading 'v' stripped
	// if applicable, which is what we want for cloud interactions.
	o.MasterKubernetesVersion = mv.String()

	if o.WorkerKubernetesVersion == "" {
		// Worker pools default to master version, which is validated by now
		o.WorkerKubernetesVersion = o.MasterKubernetesVersion
		return nil
	}
	mw, err := semver.NewVersion(o.WorkerKubernetesVersion)
	if err != nil {
		return errors.Wrap(err, "worker semver")
	}
	o.WorkerKubernetesVersion = mw.String()

	return nil
}

func (o *TemplateCreateOptions) defaultAndValidateDescription() error {
	if o.Description == "" {
		o.Description = "none"
	}
	return nil
}
