package client

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	ProxyServiceName = "Proxy"
)

type ProxyClient struct {
	*Client
}

func NewProxy() *ProxyClient {
	baseURL, _ := url.Parse(viper.GetString("proxyBaseURL"))

	return &ProxyClient{
		Client: &Client{
			service: &Service{
				name:    ProxyServiceName,
				baseURL: baseURL,
			},
		},
	}
}

func (c *ProxyClient) KubernetesGet(organizationID string, clusterID string, path string, output interface{}) error {
	token := viper.GetString("token")
	if token == "" {
		return errors.New("please provide a token")
	}

	if organizationID == "" {
		return errors.New("proxy client requires an organization ID")
	}

	if clusterID == "" {
		return errors.New("proxy client requires a cluster ID")
	}

	fullURL := fmt.Sprintf("/v3/organizations/%s/clusters/%s/k8sapi/proxy%s", organizationID, clusterID, path)
	url, _ := c.service.baseURL.Parse(fullURL)

	authHeader := fmt.Sprintf("JWT %s", token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(output).
		Get(url.String())

	if err != nil {
		return errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership %s Service responded with status %d: %s",
			c.service.name, resp.StatusCode(), resp.Body())
	}

	return nil
}
