package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/api/types"
)

var (
	userTime = "1517001176920"

	users = []types.User{
		{
			ID:      types.UUID("1234"),
			Name:    strptr("Matt K"),
			Email:   strptr("m@containership.io"),
			AddedAt: &userTime,
		},
		{
			ID:      types.UUID("4321"),
			Name:    strptr("Ashley S"),
			Email:   strptr("a@containership.io"),
			AddedAt: &userTime,
		},
	}
)

func TestNewUsers(t *testing.T) {
	a := NewUsers(nil)
	assert.NotNil(t, a)

	a = NewUsers(users)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(users))

	a = User()
	assert.NotNil(t, a)
}

func TestUsersTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewUsers(users)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(users), info.numRows)
}
