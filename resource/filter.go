package resource

// TODO make filters chainable
// TODO this is not quite the right abstraction - filter funcs?
type filterable interface {
	FilterByOwnerID(id string)
}
