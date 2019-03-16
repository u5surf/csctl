package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenCreateDefaultAndValidate(t *testing.T) {
	var opts = AccessTokenCreate{}
	// Not everything is defaultable, so empty opts is not ok
	err := opts.DefaultAndValidate()
	assert.Error(t, err, "empty opts not ok")

	opts = AccessTokenCreate{
		Name: "name",
	}
	err = opts.DefaultAndValidate()
	assert.NoError(t, err, "all required options present")

	noName := opts
	noName.Name = ""
	err = noName.DefaultAndValidate()
	assert.Error(t, err, "name is required")
}
