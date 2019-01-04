package options

import (
	"github.com/pkg/errors"
)

// ClusterCreate is the set of options required to create a cluster
type ClusterCreate struct {
	TemplateID string
	ProviderID string

	Name        string
	Environment string

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

	o.labels = map[string]string{
		"containership.io/cluster-name":        o.Name,
		"containership.io/cluster-environment": o.Environment,
	}

	return nil
}
