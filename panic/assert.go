package assert

import (
	"github.com/ymz-ncnk/assert"
	"golang.org/x/exp/constraints"
)

// Turns on assertions.
var On bool

var callback = func(msg string) { panic(msg) }

// Equal invariant will panic if two compared variables are not equal.
func Equal[T comparable](a, b T) {
	if On {
		assert.Equal(a, b, callback)
	}
}

// EqualDeep invariant will panic if two variables are not deep equal.
func EqualDeep(a, b interface{}) {
	if On {
		assert.EqualDeep(a, b, callback)
	}
}

// EqualBytes invariant will panic if two byte slices are not equal.
func EqualBytes(a, b []byte) {
	if On {
		assert.EqualBytes(a, b, callback)
	}
}

// EqualError invariant will panic if two errors are not equal.
func EqualError(a, b error) {
	if On {
		assert.EqualError(a, b, callback)
	}
}

// Bigger invariant will panic if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T) {
	if On {
		assert.Bigger(a, b, callback)
	}
}

// Lesser invariant will panic if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T) {
	if On {
		assert.Lesser(a, b, callback)
	}
}
