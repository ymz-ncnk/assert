package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type Callback func(msg string)

// Equal invariant will call callback if two compared variables are not equal.
func Equal[T comparable](a, b T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	return EqualFn(func() T { return a }, b, callback, t)
}

// EqualFn invariant will call callback if fn compared result is not equal to r.
func EqualFn[T comparable](fn func() T, r T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	a := fn()
	if a != r {
		callback(fmt.Sprintf("%v != %v", a, r))
		return false
	}
	return true
}

// EqualDeep invariant will call callback if two variables are not deep equal.
func EqualDeep(a, b interface{}, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	return EqualDeepFn(func() any { return a }, b, callback, t)
}

// EqualDeepFn invariant will call callback if fn result is not deep equal to r.
func EqualDeepFn(fn func() interface{}, r interface{},
	callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	a := fn()
	if !reflect.DeepEqual(a, r) {
		callback(fmt.Sprintf("%v != %v", a, r))
		return false
	}
	return true
}

// EqualBytes invariant will call callback if two byte slices are not equal.
func EqualBytes(a, b []byte, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	return EqualBytesFn(func() []byte { return a }, b, callback, t)
}

// EqualBytesFn invariant will call callback if fn result is not equal to r.
func EqualBytesFn(fn func() []byte, r []byte, callback Callback,
	t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	a := fn()
	if !bytes.Equal(a, r) {
		callback(fmt.Sprintf("%v != %v", a, r))
		return false
	}
	return true
}

// EqualError invariant will call callback if two errors are not equal.
func EqualError(a, b error, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	return EqualErrorFn(func() error { return a }, b, callback, t)
}

// EqualErrorFn invariant will call callback if fn result error is not equal to r.
func EqualErrorFn(fn func() error, r error, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	a := fn()
	if a == nil && r == nil {
		return true
	}
	if a != nil && r != nil {
		if !Equal(reflect.TypeOf(a), reflect.TypeOf(r), callback, t) {
			return false
		}
		if !Equal(a.Error(), r.Error(), callback, t) {
			return false
		}
		return true
	}
	callback(fmt.Sprintf("\"%v\" != \"%v\"", a, r))
	return false
}

// Bigger invariant will call callback if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	return BiggerFn(func() T { return a }, b, callback, t)
}

// BiggerFn invariant will call callback if fn result is lesser than r.
func BiggerFn[T constraints.Ordered](fn func() T, r T,
	callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	a := fn()
	if a < r {
		callback(fmt.Sprintf("%v < %v", a, r))
		return false
	}
	return true
}

// Lesser invariant will call callback if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	return LesserFn(func() T { return a }, b, callback, t)
}

// LesserFn invariant will call callback if fn result is bigger than r.
func LesserFn[T constraints.Ordered](fn func() T, r T,
	callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	a := fn()
	if a > r {
		callback(fmt.Sprintf("%v > %v", a, r))
		return false
	}
	return true
}
