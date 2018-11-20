package rest

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
)

// Interface is the set of REST actions available
type Interface interface {
	Get(path string, output interface{}) error
	Delete(path string) error
	Post(path string, body interface{}, output interface{}) error
}

// Client is a basic HTTP client for REST operations
type Client struct {
	baseURL *url.URL
	token   string
}

// Config is the configuration for a REST client
type Config struct {
	BaseURL string
	Token   string
}

// NewClient constructs a new REST client for the given config
func NewClient(cfg Config) (*Client, error) {
	// Note that this function is used over url.Parse() because it actually
	// validates the base URL format, whereas Parse() will not
	baseURL, err := url.ParseRequestURI(cfg.BaseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing base URL %q", cfg.BaseURL)
	}

	if cfg.Token == "" {
		return nil, errors.New("token is required")
	}

	return &Client{
		baseURL: baseURL,
		token:   cfg.Token,
	}, nil
}

// Get gets a resource at the given path and stores the result in output
// or returns an error
func (c *Client) Get(path string, output interface{}) error {
	url, err := c.baseURL.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "parsing path %q", path)
	}

	authHeader := fmt.Sprintf("JWT %s", c.token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(output).
		Get(url.String())

	if err != nil {
		return errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return httpErrorFromResponse(resp)
	}

	return nil
}

// Delete deletes the resource at the given path or returns an error
func (c *Client) Delete(path string) error {
	url, err := c.baseURL.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "parsing path %q", path)
	}

	authHeader := fmt.Sprintf("JWT %s", c.token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		Delete(url.String())

	if err != nil {
		return errors.Wrap(err, "error deleting resource")
	}

	if resp.IsError() {
		return httpErrorFromResponse(resp)
	}

	return nil
}

// Post posts the body to the given path and
// stores the response in output or returns an
// error
func (c *Client) Post(path string, body interface{}, output interface{}) error {
	url, err := c.baseURL.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "parsing path %q", path)
	}

	authHeader := fmt.Sprintf("JWT %s", c.token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetBody(body).
		SetResult(output).
		Post(url.String())

	if err != nil {
		return errors.Wrap(err, "error deleting resource")
	}

	if resp.IsError() {
		return httpErrorFromResponse(resp)
	}

	return nil
}

func httpErrorFromResponse(resp *resty.Response) HTTPError {
	return HTTPError{
		code:    resp.StatusCode(),
		message: string(resp.Body()),
	}
}
