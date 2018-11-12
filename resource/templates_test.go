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
		CreatedAt:     &tmplTime,
		Configuration: &tmplConfig,
	}

	tmplNoConfig = types.Template{
		ID:           types.UUID("1234"),
		ProviderName: strptr("digital_ocean"),
		OwnerID:      types.UUID("1234"),
		CreatedAt:    &tmplTime,
	}

	tmpls = []types.Template{
		tmplGood,
		{
			ID:            types.UUID("4321"),
			ProviderName:  strptr("amazon_web_services"),
			OwnerID:       types.UUID("4321"),
			CreatedAt:     &tmplTime,
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

func TestFilterByOwnerID(t *testing.T) {
	me := types.UUID("f114bfb7-0f03-497c-9522-8ab74f9fb18c")
	me1 := types.Template{
		OwnerID:      me,
		ProviderName: strptr("google"),
	}
	me2 := types.Template{
		OwnerID:      me,
		ProviderName: strptr("digital_ocean"),
	}
	notMe := types.Template{
		OwnerID:      types.UUID("00000000-0f03-497c-9522-8ab74f9fb18c"),
		ProviderName: strptr("google"),
	}
	// Not valid but shouldn't explode
	noOwner := types.Template{
		ProviderName: strptr("azure"),
	}
	tmpls := NewTemplates([]types.Template{
		me1,
		notMe,
		me2,
		noOwner,
	})

	tmpls.FilterByOwnerID(string(me))

	assert.Len(t, tmpls.items, 2)
	assert.Contains(t, tmpls.items, me1)
	assert.Contains(t, tmpls.items, me2)
	assert.NotContains(t, tmpls.items, notMe)
	assert.NotContains(t, tmpls.items, noOwner)
}

func TestFilterByEngine(t *testing.T) {
	cke1 := types.Template{
		ID:     types.UUID("1"),
		Engine: strptr(types.TemplateEngineContainershipKubernetesEngine),
	}
	cke2 := types.Template{
		ID:     types.UUID("2"),
		Engine: strptr(types.TemplateEngineContainershipKubernetesEngine),
	}
	notCKE := types.Template{
		ID:     types.UUID("3"),
		Engine: strptr("not_cke"),
	}
	nilEngine := types.Template{
		ID: types.UUID("4"),
	}
	tmpls := NewTemplates([]types.Template{
		cke1,
		notCKE,
		cke2,
		nilEngine,
	})

	tmpls.FilterByEngine("containership_kubernetes_engine")

	assert.Len(t, tmpls.items, 2)
	assert.Contains(t, tmpls.items, cke1)
	assert.Contains(t, tmpls.items, cke2)
	assert.NotContains(t, tmpls.items, notCKE)
	assert.NotContains(t, tmpls.items, nilEngine)
}
