package client

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	APIServiceName       = "API"
	ProvisionServiceName = "Provision"
)

type Client struct {
	service *Service
}

type Service struct {
	name    string
	baseURL *url.URL
}

func NewAPI() *Client {
	baseURL, _ := url.Parse(viper.GetString("apiBaseURL"))

	return &Client{
		service: &Service{
			name:    APIServiceName,
			baseURL: baseURL,
		},
	}
}

func NewProvision() *Client {
	baseURL, _ := url.Parse(viper.GetString("provisionBaseURL"))

	return &Client{
		service: &Service{
			name:    ProvisionServiceName,
			baseURL: baseURL,
		},
	}
}

func (c *Client) GetResource(path string, output interface{}) error {
	token := viper.GetString("token")
	if token == "" {
		return errors.New("please provide a token")
	}

	url, _ := c.service.baseURL.Parse(path)

	authHeader := fmt.Sprintf("JWT %s", token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(output).
		Get(url.String())

	if err != nil {
		return errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership Cloud responded with status %d: %s", resp.StatusCode(), resp.Body())
	}

	return nil
}
