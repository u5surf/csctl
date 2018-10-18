package kubeconfig

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"

	clientcmdv1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

// Config represents the minimum required fields to construct a new kubeconfig
// for accessing Containership clusters
type Config struct {
	ServerAddress string
	ClusterID     string
	UserID        string
	Token         string
}

// New constructs a new kubeconfig for the given config
func New(cfg *Config) *clientcmdv1.Config {
	clusterName := fmt.Sprintf("cs-%s", cfg.ClusterID)
	contextName := fmt.Sprintf("cs-ctx-%s", cfg.ClusterID)
	userName := fmt.Sprintf("cs-%s", cfg.UserID)

	return &clientcmdv1.Config{
		Clusters: []clientcmdv1.NamedCluster{
			{
				Name: clusterName,
				Cluster: clientcmdv1.Cluster{
					Server: cfg.ServerAddress,
				},
			},
		},

		AuthInfos: []clientcmdv1.NamedAuthInfo{
			{
				Name: userName,
				AuthInfo: clientcmdv1.AuthInfo{
					Token: cfg.Token,
				},
			},
		},

		Contexts: []clientcmdv1.NamedContext{
			{
				Name: contextName,
				Context: clientcmdv1.Context{
					Cluster:  clusterName,
					AuthInfo: userName,
				},
			},
		},

		// TODO may not want to set this by default once merging into ~/.kube/config is implemented
		CurrentContext: contextName,
	}
}

// Write writes a kubeconfig to the given writer
func Write(cfg *clientcmdv1.Config, w io.Writer) error {
	j, err := json.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "intermediate parsing to JSON")
	}

	y, err := yaml.JSONToYAML([]byte(j))
	if err != nil {
		return errors.Wrap(err, "converting to yaml")
	}

	fmt.Fprint(w, string(y))
	return nil
}
