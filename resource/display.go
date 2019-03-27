package resource

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/jsonpath"
)

// Displayable is the interface for resources that can be displayed (printed to an io.Writer)
// in various formats.
type Displayable interface {
	columns() []string

	Table(w io.Writer) error
	JSON(w io.Writer) error
	YAML(w io.Writer) error
	JSONPath(w io.Writer, template string) error
}

func assertTypes(data interface{}, listView bool) interface{} {
	switch val := data.(type) {
	case []types.Template:
		if len(val) == 1 && !listView {
			return val[1:]
		}
	}
	return data
}

func displayJSON(w io.Writer, data interface{}, listView bool) error {
	j, err := json.MarshalIndent(assertTypes(data, listView), "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshaling to JSON")
	}

	fmt.Fprintln(w, string(j))

	return nil
}

func displayYAML(w io.Writer, data interface{}, listView bool) error {
	j, err := json.Marshal(assertTypes(data, listView))
	if err != nil {
		return errors.Wrap(err, "marshaling to JSON")
	}

	y, err := yaml.JSONToYAML([]byte(j))
	if err != nil {
		return errors.Wrap(err, "converting JSON to YAML")
	}

	fmt.Fprintln(w, string(y))

	return nil
}

func displayJSONPath(w io.Writer, template string, data interface{}) error {
	jp := jsonpath.New("")
	err := jp.Parse(template)
	if err != nil {
		return errors.Wrap(err, "parsing jsonpath")
	}

	err = jp.Execute(w, data)
	if err != nil {
		return errors.Wrap(err, "executing jsonpath")
	}

	return nil
}
