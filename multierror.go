package builtinhelper

import (
	"encoding/json"
)

// MultiError represents multiple errors as an error.
type MultiError []error

// Error implements the error interface.
func (me *MultiError) Error() string {
	ss := []string{}
	for _, e := range *me {
		ss = append(ss, e.Error())
	}
	bb, _ := json.Marshal(ss)
	return string(bb)
}

// Add adds an error to MultiError.
func (me *MultiError) Add(err error) {
	*me = append(*me, err)
}
