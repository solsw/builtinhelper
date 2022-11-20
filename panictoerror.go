package builtinhelper

import (
	"errors"
)

// PanicToError converts panic to error in the following way:
//
// - if the surrounding function panics with an error, the error is recovered and returned;
//
// - if the surrounding function panics with a string, the string is recovered, wrapped into an error and returned;
//
// - otherwise the panic is reraised.
//
// PanicToError must be called from a "defer" statement:
//
// func Example() (err error) {
//   defer func() {
//     PanicToError(recover(), &err)
//   }()
//   return nil
// }
//
// (See tests.)
func PanicToError(panicArg interface{}, err *error) {
	if panicArg == nil {
		return
	}
	if err == nil {
		panic(panicArg)
	}
	e, isError := panicArg.(error)
	if isError {
		*err = e
		return
	}
	s, isString := panicArg.(string)
	if isString {
		*err = errors.New(s)
		return
	}
	panic(panicArg)
}
