package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type validateScaleTest struct {
	from int32
	up   int32
	down int32
	to   int32

	shouldError bool
	expected    int32
	msg         string
}

var validateScaleTests = []validateScaleTest{
	{
		from: 1,
		up:   countUnset,
		down: 1,
		to:   countUnset,

		shouldError: false,
		expected:    0,
		msg:         "scale down to 0 is allowed",
	},
	{
		from: 2,
		up:   countUnset,
		down: countUnset,
		to:   0,

		shouldError: false,
		expected:    0,
		msg:         "scale target 0 is allowed",
	},
	{
		from: 2,
		up:   -2,
		down: countUnset,
		to:   countUnset,

		shouldError: true,
		msg:         "scale up by less than 1",
	},
	{
		from: 2,
		up:   countUnset,
		down: -2,
		to:   countUnset,

		shouldError: true,
		msg:         "scale down by less than 1",
	},
	{
		from: 2,
		up:   countUnset,
		down: 3,
		to:   countUnset,

		shouldError: true,
		msg:         "scale down below 0",
	},
	{
		from: 2,
		up:   countUnset,
		down: countUnset,
		to:   -1,

		shouldError: true,
		msg:         "scale target below 0",
	},
	{
		from: 2,
		up:   3,
		down: countUnset,
		to:   countUnset,

		shouldError: false,
		expected:    5,
		msg:         "scale up by N",
	},
	{
		from: 5,
		up:   countUnset,
		down: 2,
		to:   countUnset,

		shouldError: false,
		expected:    3,
		msg:         "scale down by N",
	},
	{
		from: 5,
		up:   countUnset,
		down: countUnset,
		to:   2,

		shouldError: false,
		expected:    2,
		msg:         "scale to target",
	},
}

func TestValidateAndGetTargetCount(t *testing.T) {
	for _, c := range validateScaleTests {
		target, err := validateAndGetTargetCount(c.from, c.up, c.down, c.to)
		if c.shouldError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expected, target)
		}
	}
}

func TestExactlyOneSet(t *testing.T) {
	assert.False(t, exactlyOneSet(countUnset, countUnset, countUnset))

	assert.True(t, exactlyOneSet(2, countUnset, countUnset))
	assert.True(t, exactlyOneSet(countUnset, 3, countUnset))
	assert.True(t, exactlyOneSet(countUnset, countUnset, 4))

	assert.False(t, exactlyOneSet(2, countUnset, 4))
	assert.False(t, exactlyOneSet(countUnset, 3, 4))
	assert.False(t, exactlyOneSet(2, 3, countUnset))
	assert.False(t, exactlyOneSet(2, 3, 4))
}
