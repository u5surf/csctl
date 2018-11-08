package resource

type resourceInterface interface {
	Displayable

	Name() string
	Plural() string
	Aliases() []string
	HasAlias(name string) bool
}

type resource struct {
	resourceInterface

	name    string
	plural  string
	aliases []string
}

func (r *resource) Name() string {
	return r.name
}

func (r *resource) Plural() string {
	return r.plural
}

func (r *resource) Aliases() []string {
	return append(r.aliases, r.name, r.plural)
}

func (r *resource) HasAlias(name string) bool {
	if name == r.name {
		return true
	}

	if r.plural != "" && name == r.plural {
		return true
	}

	for _, a := range r.aliases {
		if name == a {
			return true
		}
	}

	return false
}
