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
	providersSingle = []types.Provider{
		{
			ID:          types.UUID("1234"),
			Provider:    strptr("google"),
			Description: strptr("google description"),
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

func TestProvidersDisableListView(t *testing.T) {
	a := NewProviders(providersSingle)
	assert.NotNil(t, a)
	a.resource.DisableListView()
	assert.Equal(t, a.resource.listView, false)
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

func TestProvidersJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewProviders(providersSingle)
	err := a.JSON(buf)
	assert.Nil(t, err)
	a.resource.DisableListView()
	err = a.JSON(buf)
	assert.Nil(t, err)
}

func TestProvidersYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewProviders(providersSingle)
	err := a.YAML(buf)
	assert.Nil(t, err)
	a.resource.DisableListView()
	err = a.YAML(buf)
	assert.Nil(t, err)
}
