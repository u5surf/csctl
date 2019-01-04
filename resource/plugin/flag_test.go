package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "Unknown", (TypeCNI - 1).String())
	assert.Equal(t, "Unknown", (TypeLogs + 1).String())
	assert.Equal(t, "Logs", TypeLogs.String())
}

func TestParse(t *testing.T) {
	f := Flag{}
	impl, version, err := f.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "", impl)
	assert.Equal(t, "", version)

	f = Flag{Val: "calico@3.3.0"}
	impl, version, err = f.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "calico", impl)
	assert.Equal(t, "3.3.0", version)

	// Leading 'v' is ok
	f = Flag{Val: "calico@v3.3.0"}
	impl, version, err = f.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "calico", impl)
	assert.Equal(t, "3.3.0", version)

	f = Flag{Val: "calico"}
	impl, version, err = f.Parse()
	assert.NoError(t, err)
	assert.Equal(t, "calico", impl)
	assert.Equal(t, "", version)

	f = Flag{Val: "calico@"}
	impl, version, err = f.Parse()
	assert.Error(t, err, "@ but no semver")

	f = Flag{Val: "calico@123asdf"}
	impl, version, err = f.Parse()
	assert.Error(t, err, "invalid semver")
}
