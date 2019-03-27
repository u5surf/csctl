package resource

const (
	emptyColState = "<none>"
)

type resourceInterface interface {
	Displayable

	Name() string
	Plural() string
	Aliases() []string
	HasAlias(name string) bool
	DisableListView()
}

type resource struct {
	resourceInterface

	name     string
	plural   string
	aliases  []string
	listView bool
}

func (r *resource) Name() string {
	return r.name
}

func (r *resource) Plural() string {
	return r.plural
}

func (r *resource) Aliases() []string {
	if r.plural != "" {
		return append(r.aliases, r.plural)
	}

	return r.aliases
}

func (r *resource) DisableListView() {
	r.listView = false
}
