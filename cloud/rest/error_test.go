package rest

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	emptyError = HTTPError{}

	codeOnlyError = HTTPError{
		code: 500,
	}

	internalServerError = HTTPError{
		code:    500,
		message: "uh oh",
	}

	notFoundError = HTTPError{
		code:    404,
		message: "not found",
	}
)

func TestError(t *testing.T) {
	s := emptyError.Error()
	assert.NotEmpty(t, s, "error string never empty")

	// This is kinda silly but works for now
	s = codeOnlyError.Error()
	assert.NotEmpty(t, s, "correct string if not message")
	assert.Contains(t, s, strconv.Itoa(codeOnlyError.code), "contains code")
	assert.NotContains(t, s, ":", "doesn't contain colon")

	s = internalServerError.Error()
	assert.Contains(t, s, strconv.Itoa(internalServerError.code), "contains code")
	assert.Contains(t, s, internalServerError.message, "contains message when available")
}

func TestIsNotFound(t *testing.T) {
	isNotFound := notFoundError.IsNotFound()
	assert.True(t, isNotFound, "not found error is 'not found'")

	isNotFound = internalServerError.IsNotFound()
	assert.False(t, isNotFound, "internal server error is not 'not found'")
}
