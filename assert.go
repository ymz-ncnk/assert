package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

type Callback func(msg string)

// Equal invariant will call callback if two compared variables are not equal.
func Equal[T comparable](a, b T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if a != b {
		callback(fmt.Sprintf("%v != %v", a, b))
		return false
	}
	return true
}

// EqualDeep invariant will call callback if two variables are not deep equal.
func EqualDeep(a, b any, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if !reflect.DeepEqual(a, b) {
		callback(fmt.Sprintf("%v != %v", a, b))
		return false
	}
	return true
}

// EqualBytes invariant will call callback if two byte slices are not equal.
func EqualBytes(a, b []byte, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if !bytes.Equal(a, b) {
		callback(fmt.Sprintf("%v != %v", a, b))
		return false
	}
	return true
}

// EqualError invariant will call callback if two errors are not equal.
func EqualError(a, b error, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		if !Equal(reflect.TypeOf(a), reflect.TypeOf(b), callback, t) {
			return false
		}
		if !Equal(a.Error(), b.Error(), callback, t) {
			return false
		}
		return true
	}
	callback(fmt.Sprintf("\"%v\" != \"%v\"", a, b))
	return false
}

// Bigger invariant will call callback if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if a < b {
		callback(fmt.Sprintf("%v < %v", a, b))
		return false
	}
	return true
}

// Lesser invariant will call callback if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T, callback Callback, t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if a > b {
		callback(fmt.Sprintf("%v > %v", a, b))
		return false
	}
	return true
}

func SameTime(a, b time.Time, delta time.Duration, callback Callback,
	t *testing.T) bool {
	if t != nil {
		t.Helper()
	}
	if !(b.Before(a.Add(delta)) && a.Before(b.Add(delta))) {
		callback(fmt.Sprintf("%v != %v", a, b))
		return false
	}
	return true
}
