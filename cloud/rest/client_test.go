package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// Dummy JWT from jwt.io. Valid format but probably expired :)
	validJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	// Valid base URL
	containershipURL = "https://api.containership.io"
)

// TODO tests for REST actions

func TestNewClient(t *testing.T) {
	// Good config
	client, err := NewClient(&Config{
		BaseURL: containershipURL,
		Token:   validJWT,
	})
	assert.Nil(t, err)
	assert.NotNil(t, client)

	// BaseURL is required
	client, err = NewClient(&Config{
		BaseURL: "",
		Token:   validJWT,
	})
	assert.NotNil(t, err)

	// Invalid BaseURL
	client, err = NewClient(&Config{
		BaseURL: "invalid",
		Token:   validJWT,
	})
	assert.NotNil(t, err)

	// Token is required
	client, err = NewClient(&Config{
		BaseURL: containershipURL,
		Token:   "",
	})
	assert.NotNil(t, err)
}
