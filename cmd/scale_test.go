package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExactlyOneSet(t *testing.T) {
	assert.False(t, exactlyOneSet(0, 0, 0))

	assert.True(t, exactlyOneSet(2, 0, 0))
	assert.True(t, exactlyOneSet(0, 3, 0))
	assert.True(t, exactlyOneSet(0, 0, 4))

	assert.False(t, exactlyOneSet(2, 0, 4))
	assert.False(t, exactlyOneSet(0, 3, 4))
	assert.False(t, exactlyOneSet(2, 3, 0))
	assert.False(t, exactlyOneSet(2, 3, 4))
}
