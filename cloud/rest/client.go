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
	Patch(path string, body interface{}, output interface{}) error
}

// Client is a basic HTTP client for REST operations
type Client struct {
	baseURL      *url.URL
	token        string
	debugEnabled bool
}

// Config is the configuration for a REST client
type Config struct {
	BaseURL      string
	Token        string
	DebugEnabled bool
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
		baseURL:      baseURL,
		token:        cfg.Token,
		debugEnabled: cfg.DebugEnabled,
	}, nil
}

// Get gets a resource at the given path and stores the result in output
// or returns an error
func (c *Client) Get(path string, output interface{}) error {
	return c.execute(resty.MethodGet, path, nil, output)
}

// Delete deletes the resource at the given path or returns an error
func (c *Client) Delete(path string) error {
	return c.execute(resty.MethodDelete, path, nil, nil)
}

// Post posts the body to the given path and stores the response in output or
// returns an error
func (c *Client) Post(path string, body interface{}, output interface{}) error {
	return c.execute(resty.MethodPost, path, body, output)
}

// Patch patches the body to the given path and stores the response in output or
// returns an error
func (c *Client) Patch(path string, body interface{}, output interface{}) error {
	return c.execute(resty.MethodPatch, path, body, output)
}

func (c *Client) execute(verb string, path string, body interface{}, output interface{}) error {
	url, err := c.baseURL.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "parsing path %q", path)
	}

	authHeader := fmt.Sprintf("JWT %s", c.token)

	req := resty.SetDebug(c.debugEnabled).R().SetHeader("Authorization", authHeader)

	if body != nil {
		req.SetBody(body)
	}

	if output != nil {
		req.SetResult(output)
	}

	resp, err := req.Execute(verb, url.String())

	if err != nil {
		return errors.Wrapf(err, "error %sing resource", verb)
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
