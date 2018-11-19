package rest

import (
	"fmt"
)

// HTTPError implements the error interface for an HTTP-related error
type HTTPError struct {
	code    int
	message string
}

// Error returns a string representing this error
func (e HTTPError) Error() string {
	s := fmt.Sprintf("server responded with status %d", e.code)
	if e.message != "" {
		s = fmt.Sprintf("%s: %s", s, e.message)
	}

	return s
}

// IsNotFound returns true if this is a "not found" error, else false
func (e HTTPError) IsNotFound() bool {
	return e.code == 404
}
