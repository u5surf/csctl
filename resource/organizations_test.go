package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	org1    = "test1"
	org2    = "test2"
	orgTime = "1517001176920"

	orgs = []types.Organization{
		{
			Name:      &org1,
			ID:        types.UUID("1234"),
			OwnerID:   types.UUID("1234"),
			CreatedAt: &orgTime,
		},
		{
			Name:      &org2,
			ID:        types.UUID("4321"),
			OwnerID:   types.UUID("4321"),
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
