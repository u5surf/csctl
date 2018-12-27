package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	regTime = "1517001176920"

	regs = []types.Registry{
		{
			ID:             types.UUID("1234"),
			CreatedAt:      &regTime,
			Description:    strptr("logs"),
			Provider:       strptr("kubernetes"),
			Serveraddress:  strptr("v1.0.0"),
		},
		{
			ID:             types.UUID("4321"),
			CreatedAt:      &regTime,
			Description:    strptr("metrics"),
			Provider:       strptr("prometheus"),
			Serveraddress:  strptr("2.0.0"),
		},
	}
)

func TestNewRegistries(t *testing.T) {
	a := NewRegistries(nil)
	assert.NotNil(t, a)

	a = NewRegistries(regs)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(regs))

	a = Registry()
	assert.NotNil(t, a)
}

func TestRegistriesTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewRegistries(regs)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(regs), info.numRows)
}
