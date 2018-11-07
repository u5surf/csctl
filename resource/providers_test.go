package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	providerTime = "1517001176920"

	providers = []types.Provider{
		{
			ID:          types.UUID("1234"),
			Provider:    strptr("google"),
			Description: strptr("google description"),
			CreatedAt:   &providerTime,
		},
		{
			ID:          types.UUID("4321"),
			Provider:    strptr("amazon_web_services"),
			Description: strptr("aws description"),
			CreatedAt:   &providerTime,
		},
	}
)

func TestNewProviders(t *testing.T) {
	a := NewProviders(nil)
	assert.NotNil(t, a)

	a = NewProviders(providers)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(providers))

	a = Provider()
	assert.NotNil(t, a)
}

func TestProvidersTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewProviders(providers)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(providers), info.numRows)
}
