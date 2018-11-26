package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

const (
	// Dummy JWT from jwt.io. Valid format but probably expired :)
	validJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	// Valid base URL
	containershipURL = "https://api.containership.io"

	// Valid path
	testPath = "/v3/organizations/a5df7ade-f185-40a0-a379-0324af86b491"
)

// TODO tests for REST actions

func TestNewClient(t *testing.T) {
	// Good config
	client, err := NewClient(Config{
		BaseURL: containershipURL,
		Token:   validJWT,
	})
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// BaseURL is required
	client, err = NewClient(Config{
		BaseURL: "",
		Token:   validJWT,
	})
	assert.NotNil(t, err)

	// Invalid BaseURL
	client, err = NewClient(Config{
		BaseURL: "invalid",
		Token:   validJWT,
	})
	assert.NotNil(t, err)

	// Token is required
	client, err = NewClient(Config{
		BaseURL: containershipURL,
		Token:   "",
	})
	assert.NotNil(t, err)
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method, "correct HTTP method used")
	}))
	defer ts.Close()

	client, err := NewClient(Config{
		BaseURL: ts.URL,
		Token:   validJWT,
	})
	assert.NoError(t, err)

	var out types.Organization
	err = client.Get(testPath, &out)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "DELETE", r.Method, "correct HTTP method used")
	}))
	defer ts.Close()

	client, err := NewClient(Config{
		BaseURL: ts.URL,
		Token:   validJWT,
	})
	assert.NoError(t, err)

	err = client.Delete(testPath)
	assert.NoError(t, err)
}

func TestPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method, "correct HTTP method used")
	}))
	defer ts.Close()

	client, err := NewClient(Config{
		BaseURL: ts.URL,
		Token:   validJWT,
	})
	assert.NoError(t, err)

	var body, out interface{}
	err = client.Post(testPath, body, out)
	assert.NoError(t, err)
}

func TestPatch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PATCH", r.Method, "correct HTTP method used")
	}))
	defer ts.Close()

	client, err := NewClient(Config{
		BaseURL: ts.URL,
		Token:   validJWT,
	})
	assert.NoError(t, err)

	var body, out interface{}
	err = client.Patch(testPath, body, out)
	assert.NoError(t, err)
}

func testServerWithStatus(status int) (*httptest.Server, *Client, error) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))

	client, err := NewClient(Config{
		BaseURL: ts.URL,
		Token:   validJWT,
	})
	if err != nil {
		return nil, nil, err
	}

	return ts, client, nil
}

func testServerNonEmptyBodyChecker(t *testing.T) (*httptest.Server, *Client, error) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err, "body is readable")
		assert.NotEmpty(t, body, "non-empty body")
	}))

	client, err := NewClient(Config{
		BaseURL: ts.URL,
		Token:   validJWT,
	})
	if err != nil {
		return nil, nil, err
	}

	return ts, client, nil
}

func TestExecute(t *testing.T) {
	ts, client, err := testServerWithStatus(http.StatusOK)
	assert.NoError(t, err)
	defer ts.Close()

	err = client.Get("%%", nil)
	assert.Error(t, err, "bad path results in error")

	ts, client, err = testServerWithStatus(http.StatusInternalServerError)
	assert.NoError(t, err)
	defer ts.Close()

	var out types.Organization
	err = client.Get(testPath, &out)
	assert.Error(t, err, "server error results in client error")

	ts, client, err = testServerWithStatus(http.StatusNotFound)
	assert.NoError(t, err)
	defer ts.Close()

	err = client.Get(testPath, &out)
	assert.Error(t, err, "not found error results in client error")
	nfErr, ok := err.(HTTPError)
	assert.True(t, ok, "not found error results in HTTPError type")
	assert.True(t, nfErr.IsNotFound(), "not found error results in 'not found' error type")

	ts, client, err = testServerNonEmptyBodyChecker(t)
	assert.NoError(t, err)
	defer ts.Close()
	body := types.Organization{
		ID: "a5df7ade-f185-40a0-a379-0324af86b491",
	}
	err = client.Post(testPath, &body, &out)
}

func TestHttpErrorFromResponse(t *testing.T) {
	resp := &resty.Response{
		RawResponse: &http.Response{
			StatusCode: 404,
		},
	}

	err := httpErrorFromResponse(resp)
	assert.Error(t, err)
	assert.True(t, err.IsNotFound())
}
