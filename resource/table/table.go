package table

import (
	"fmt"
	"io"
	"strings"

	"text/tabwriter"
)

type Table struct {
	*tabwriter.Writer
}

func New(w io.Writer, header []string) *Table {
	tw := new(tabwriter.Writer)

	tw.Init(w, 2, 8, 2, '\t', 0)

	var headerTransformed = make([]string, len(header))
	for i, h := range header {
		headerTransformed[i] = strings.ToUpper(h)
	}
	fmt.Fprintln(tw, strings.Join(headerTransformed, "\t"))

	return &Table{tw}
}

func (t *Table) Append(row []string) {
	fmt.Fprintln(t, strings.Join(row, "\t"))
}

func (t *Table) Render() {
	t.Flush()
}
