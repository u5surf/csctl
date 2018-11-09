package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	r := resource{
		name: "plugin",
	}

	assert.Equal(t, r.name, r.Name())
}

func TestPlural(t *testing.T) {
	r := resource{
		name:   "plugin",
		plural: "plugins",
	}

	assert.Equal(t, r.plural, r.Plural())
}

func TestAliases(t *testing.T) {
	r := resource{
		name:    "plugin",
		plural:  "plugins",
		aliases: []string{"plug", "plugs"},
	}

	assert.Contains(t, r.Aliases(), r.plural)
	assert.Len(t, r.Aliases(), len(r.aliases)+1, "aliases includes plural")

	r.plural = ""
	assert.Len(t, r.Aliases(), len(r.aliases), "no plural")

	r.aliases = []string{}
	assert.Empty(t, r.Aliases(), []string{}, "no plural or aliases")

	r.plural = "plugins"
	assert.Equal(t, r.Aliases(), []string{r.plural}, "plural only")
}
