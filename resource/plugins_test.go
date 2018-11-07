package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	plugTime = "1517001176920"

	plugs = []types.Plugin{
		{
			ID:             types.UUID("1234"),
			Type:           strptr("logs"),
			Implementation: strptr("kubernetes"),
			Version:        strptr("v1.0.0"),
			CreatedAt:      &plugTime,
		},
		{
			ID:             types.UUID("4321"),
			Type:           strptr("metrics"),
			Implementation: strptr("prometheus"),
			Version:        strptr("2.0.0"),
			CreatedAt:      &plugTime,
		},
	}
)

func TestNewPlugins(t *testing.T) {
	a := NewPlugins(nil)
	assert.NotNil(t, a)

	a = NewPlugins(plugs)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(plugs))

	a = Plugin()
	assert.NotNil(t, a)
}

func TestPluginsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewPlugins(plugs)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(plugs), info.numRows)
}
