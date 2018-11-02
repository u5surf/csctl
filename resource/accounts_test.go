package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	acct1    = "test1"
	acct2    = "test2"
	acctTime = "1517001176920"

	accts = []types.Account{
		{
			Name:      &acct1,
			ID:        types.UUID("1234"),
			CreatedAt: &acctTime,
		},
		{
			Name:      &acct2,
			ID:        types.UUID("4321"),
			CreatedAt: &acctTime,
		},
	}
)

func TestNewAccounts(t *testing.T) {
	a := NewAccounts(nil)
	assert.NotNil(t, a)

	a = NewAccounts(accts)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(accts))

	a = Account()
	assert.NotNil(t, a)
}

func TestAccountsTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAccounts(accts)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(accts), info.numRows)
}
