package assert

import (
	"time"

	"github.com/ymz-ncnk/assert"
	"golang.org/x/exp/constraints"
)

// Turns on assertions.
var On bool

func newCallback(msgAndArgs ...any) assert.Callback {
	return func(msg string) {
		panic(assert.Format(msg, msgAndArgs...))
	}
}

// Equal invariant will panic if two compared variables are not equal.
func Equal[T comparable](a, b T, msgAndArgs ...any) {
	if On {
		assert.Equal(a, b, newCallback(msgAndArgs...), nil)
	}
}

// EqualDeep invariant will panic if two variables are not deep equal.
func EqualDeep(a, b any, msgAndArgs ...any) {
	if On {
		assert.EqualDeep(a, b, newCallback(msgAndArgs...), nil)
	}
}

// EqualBytes invariant will panic if two byte slices are not equal.
func EqualBytes(a, b []byte, msgAndArgs ...any) {
	if On {
		assert.EqualBytes(a, b, newCallback(msgAndArgs...), nil)
	}
}

// EqualError invariant will panic if two errors are not equal.
func EqualError(a, b error, msgAndArgs ...any) {
	if On {
		assert.EqualError(a, b, newCallback(msgAndArgs...), nil)
	}
}

// Bigger invariant will panic if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T, msgAndArgs ...any) {
	if On {
		assert.Bigger(a, b, newCallback(msgAndArgs...), nil)
	}
}

// Lesser invariant will panic if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T, msgAndArgs ...any) {
	if On {
		assert.Lesser(a, b, newCallback(msgAndArgs...), nil)
	}
}

// SameTime invariant will panic if two times are not equal.
func SameTime(a, b time.Time, delta time.Duration, msgAndArgs ...any) {
	if On {
		assert.SameTime(a, b, delta, newCallback(msgAndArgs...), nil)
	}
}
