package builtinhelper

import (
	"errors"
	"testing"

	"github.com/solsw/oshelper"
)

func TestMultiError_Error(t *testing.T) {
	me1 := &MultiError{}
	me1.Add(errors.New("error1"))
	me2 := &MultiError{}
	me2.Add(errors.New("error1"))
	me2.Add(errors.New("error2"))
	tests := []struct {
		name string
		me   *MultiError
		want string
	}{
		{name: "0", me: &MultiError{}, want: ""},
		{name: "1", me: me1, want: "error1"},
		{name: "2", me: me2, want: "error1" + oshelper.NewLine + "error2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.me.Error(); got != tt.want {
				t.Errorf("MultiError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
