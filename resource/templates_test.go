package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	tmplTime       = "1517001176920"
	tmplK8sVersion = "1.12.1"

	tmplConfig = types.TemplateConfiguration{
		Variable: types.TemplateVariableMap{
			"np0": types.TemplateVariableDefault{
				Default: &types.TemplateNodePool{
					Count:             int32ptr(3),
					Etcd:              true,
					IsSchedulable:     true,
					KubernetesMode:    strptr("master"),
					KubernetesVersion: &tmplK8sVersion,
					Name:              strptr("master-pool"),
					Type:              strptr("node_pool"),
				},
			},
		},
	}

	tmplGood = types.Template{
		ID:            types.UUID("1234"),
		ProviderName:  strptr("google"),
		OwnerID:       types.UUID("1234"),
		CreatedAt:     tmplTime,
		Configuration: &tmplConfig,
	}

	tmplNoConfig = types.Template{
		ID:           types.UUID("1234"),
		ProviderName: strptr("digital_ocean"),
		OwnerID:      types.UUID("1234"),
		CreatedAt:    tmplTime,
	}

	tmpls = []types.Template{
		tmplGood,
		{
			ID:            types.UUID("4321"),
			ProviderName:  strptr("amazon_web_services"),
			OwnerID:       types.UUID("4321"),
			CreatedAt:     tmplTime,
			Configuration: &tmplConfig,
		},
	}
)

func TestNewTemplates(t *testing.T) {
	a := NewTemplates(nil)
	assert.NotNil(t, a)

	a = NewTemplates(tmpls)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(tmpls))

	a = Template()
	assert.NotNil(t, a)
}

func TestTemplatesTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewTemplates(tmpls)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(tmpls), info.numRows)
}

func TestGetKubernetesMasterVersion(t *testing.T) {
	version, err := getMasterKubernetesVersion(&tmplGood)
	assert.Nil(t, err)
	assert.Equal(t, version, tmplK8sVersion)

	version, err = getMasterKubernetesVersion(&tmplNoConfig)
	assert.NotNil(t, err)
}
