package resource

import (
	"encoding/json"
	"fmt"
	"io"

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

func displayJSON(w io.Writer, data interface{}) error {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshaling to JSON")
	}

	fmt.Fprintln(w, string(j))

	return nil
}

func displayYAML(w io.Writer, data interface{}) error {
	j, err := json.Marshal(data)
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
