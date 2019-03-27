package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	orgTime = "1517001176920"

	orgs = []types.Organization{
		{
			Name:      strptr("test1"),
			ID:        types.UUID("1234"),
			OwnerID:   types.UUID("1234"),
			CreatedAt: &orgTime,
		},
		{
			Name:      strptr("test2"),
			ID:        types.UUID("4321"),
			OwnerID:   types.UUID("4321"),
			CreatedAt: &orgTime,
		},
	}
	orgsSingle = []types.Organization{
		{
			Name:      strptr("test3"),
			ID:        types.UUID("1234"),
			OwnerID:   types.UUID("1234"),
			CreatedAt: &orgTime,
		},
	}
)

func TestNewOrganizations(t *testing.T) {
	a := NewOrganizations(nil)
	assert.NotNil(t, a)

	a = NewOrganizations(orgs)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(orgs))

	a = Organization()
	assert.NotNil(t, a)
}

func TestOrganizationsDisableListView(t *testing.T) {
	a := NewOrganizations(orgsSingle)
	assert.NotNil(t, a)
	a.resource.DisableListView()
	assert.Equal(t, a.resource.listView, false)
}

func TestOrganizationsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewOrganizations(orgs)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(orgs), info.numRows)
}

func TestOrganizationsJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewOrganizations(orgsSingle)
	err := a.JSON(buf)
	assert.Nil(t, err)
	a.resource.DisableListView()
	err = a.JSON(buf)
	assert.Nil(t, err)
}

func TestOrganizationsYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewOrganizations(orgsSingle)
	err := a.YAML(buf)
	assert.Nil(t, err)
	a.resource.DisableListView()
	err = a.YAML(buf)
	assert.Nil(t, err)
}
