package resource

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type tableInfo struct {
	numHeaderCols int
	numCols       int
	numRows       int // excluding header
}

// getTableInfo provides just enough info about a printed table for sanity
// checking.
// Some assumptions are made that may have to evolve in the future.
func getTableInfo(buf *bytes.Buffer) (*tableInfo, error) {
	header, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(header)
	numHeaderCols := len(fields)

	// offset is advanced so this ignores header
	numCols, numRows, err := countColumnsAndRows(buf)
	if err != nil {
		return nil, err
	}

	return &tableInfo{
		numHeaderCols: numHeaderCols,
		numCols:       numCols,
		numRows:       numRows,
	}, nil
}

// countColumnsAndRows returns the number of columns and rows
// in the non-header portion of the table, or an error if parsing fails.
func countColumnsAndRows(buf *bytes.Buffer) (int, int, error) {
	numCols := -1
	numRows := 0
	var line string
	var err error
	for err == nil {
		line, err = buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, err
		}

		thisRowCols := len(regexp.MustCompile("\t+").Split(line, -1))
		if numRows == 0 {
			numCols = thisRowCols
		}

		if thisRowCols != numCols {
			return 0, 0, errors.Errorf("inconsistent column count: expected %d, got %d at row %d", numCols, thisRowCols, numRows)
		}

		numRows++
	}

	return numCols, numRows, nil
}
