package builtinhelper

import (
	"errors"
	"testing"
)

func Func1() (res string, err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	return "qwerty", nil
}

func TestPanicToError1(t *testing.T) {
	_, err := Func1()
	if err != nil {
		t.Errorf("PanicToError1: err != nil")
	}
}

var ErrFunc2 = errors.New("Func2")

func Func2() (res int, err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	panic(ErrFunc2)
}

func TestPanicToError2(t *testing.T) {
	_, err := Func2()
	if err != ErrFunc2 {
		t.Errorf("PanicToError2: err != ErrFunc2")
	}
}

func Func3() (res bool, err error) {
	defer func() {
		PanicToError(recover(), &err)
	}()
	panic("Func3")
}

func TestPanicToError3(t *testing.T) {
	_, err := Func3()
	if err.Error() != "Func3" {
		t.Errorf(`PanicToError3: err != "Func3"`)
	}
}
