package api

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type APIInterface interface {
	OrganizationsGetter
	AccountGetter
}

type APIClient struct {
	name string
	// TODO make REST client and remove baseURL below
	// restClient RESTClient
	baseURL *url.URL
	token   string
}

type Config struct {
	Token   string
	BaseURL url.URL
}

// TODO don't depend on env / viper here
// Semi-weird because config comes through clientset
func New(cfg *Config) (*APIClient, error) {
	baseURL, err := url.Parse(viper.GetString("apiBaseURL"))
	if err != nil {
		return nil, errors.Wrap(err, "parsing apiBaseURL")
	}

	return &APIClient{
		name:    "API",
		baseURL: baseURL,
		token:   cfg.Token,
	}, nil
}

func (c *APIClient) Organizations() OrganizationInterface {
	return newOrganizations(c)
}

func (c *APIClient) Account() AccountInterface {
	return newAccount(c)
}

// TODO RESTClient
func (c *APIClient) GetResource(path string, output interface{}) error {
	url, _ := c.baseURL.Parse(path)

	authHeader := fmt.Sprintf("JWT %s", c.token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(output).
		Get(url.String())

	if err != nil {
		return errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership %s Service responded with status %d: %s",
			c.name, resp.StatusCode(), resp.Body())
	}

	return nil
}

func (c *APIClient) DeleteResource(path string) error {
	url, _ := c.baseURL.Parse(path)

	authHeader := fmt.Sprintf("JWT %s", c.token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		Delete(url.String())

	if err != nil {
		return errors.Wrap(err, "error deleting resource")
	}

	if resp.IsError() {
		return errors.Errorf("Containership Cloud responded with status %d: %s", resp.StatusCode(), resp.Body())
	}

	return nil
}
