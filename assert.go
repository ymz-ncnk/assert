package assert

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Turns on assertions.
var On bool

// Equal invariant will panic if two compared variables are not equal.
func Equal[T comparable](a, b T) {
	EqualFn(func() T { return a }, b)
}

// EqualFn invariant will panic if fn compared result is not equal to r.
func EqualFn[T comparable](fn func() T, r T) {
	a := fn()
	if On && a != r {
		panic(fmt.Sprintf("%v != %v", a, r))
	}
}

// EqualDeep invariant will panic if two variables are not deep equal.
func EqualDeep(a, b interface{}) {
	EqualDeepFn(func() interface{} { return a }, b)
}

// EqualDeepFn invariant will panic if fn result is not deep equal to r.
func EqualDeepFn(fn func() interface{}, r interface{}) {
	a := fn()
	if On && !reflect.DeepEqual(a, r) {
		panic(fmt.Sprintf("%v != %v", a, r))
	}
}

// EqualError invariant will panic if two errors are not equal.
func EqualError(a, b error) {
	EqualErrorFn(func() error { return a }, b)
}

// EqualFn invariant will panic if fn result error is not equal to r.
func EqualErrorFn(fn func() error, r error) {
	if !On {
		return
	}
	a := fn()
	if a == nil && r == nil {
		return
	}
	if a != nil && r != nil {
		Equal(reflect.TypeOf(a), reflect.TypeOf(r))
		Equal(a.Error(), r.Error())
		return
	}
	panic(fmt.Sprintf("\"%v\" != \"%v\"", a, r))
}

// Bigger invariant will panic if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T) {
	BiggerFn(func() T { return a }, b)
}

// EqualFn invariant will panic if fn result is lesser than r.
func BiggerFn[T constraints.Ordered](fn func() T, r T) {
	a := fn()
	if On && a < r {
		panic(fmt.Sprintf("%v < %v", a, r))
	}
}

// Lesser invariant will panic if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T) {
	LesserFn(func() T { return a }, b)
}

// EqualFn invariant will panic if fn result is bigger than r.
func LesserFn[T constraints.Ordered](fn func() T, r T) {
	a := fn()
	if On && a > r {
		panic(fmt.Sprintf("%v > %v", a, r))
	}
}
