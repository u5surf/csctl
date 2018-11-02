package table

import (
	"fmt"
	"io"
	"strings"

	"text/tabwriter"
)

// Table is an easier-to-use tabwriter with an Append interface
type Table struct {
	*tabwriter.Writer
}

// New constructs a new Table with the given header for writing to the given
// io.Writer
func New(w io.Writer, header []string) *Table {
	tw := new(tabwriter.Writer)

	tw.Init(w, 2, 8, 2, '\t', 0)

	var headerTransformed = make([]string, len(header))
	for i, h := range header {
		h = strings.ToUpper(h)
		headerTransformed[i] = strings.Replace(h, " ", "_", -1)
	}
	fmt.Fprintln(tw, strings.Join(headerTransformed, "\t"))

	return &Table{tw}
}

// Append appends a complete row to the Table
func (t *Table) Append(row []string) {
	fmt.Fprintln(t, strings.Join(row, "\t"))
}

// Render flushes the table to the io.Writer
func (t *Table) Render() {
	t.Flush()
}
