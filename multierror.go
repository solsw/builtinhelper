package builtinhelper

import (
	"strings"

	"github.com/solsw/oshelper"
)

// MultiError represents multiple errors as an error.
type MultiError []error

// Error implements the error interface.
func (me *MultiError) Error() string {
	var b strings.Builder
	for _, e := range *me {
		if b.Len() > 0 {
			b.WriteString(oshelper.NewLine)
		}
		b.WriteString(e.Error())
	}
	return b.String()
}

// Add adds an error to MultiError.
func (me *MultiError) Add(err error) {
	*me = append(*me, err)
}
