package assert

import (
	"testing"

	"github.com/ymz-ncnk/assert"
	"golang.org/x/exp/constraints"
)

func newCallback(t *testing.T) assert.Callback {
	return func(msg string) { t.Error(msg) }
}

// Equal invariant will call t.Error() if two compared variables are not equal.
func Equal[T comparable](a, b T, t *testing.T) {
	assert.Equal(a, b, newCallback(t))
}

// EqualDeep invariant will call t.Error() if two variables are not deep equal.
func EqualDeep(a, b interface{}, t *testing.T) {
	assert.EqualDeep(a, b, newCallback(t))
}

// EqualBytes invariant will call t.Error() if two byte slices are not equal.
func EqualBytes(a, b []byte, t *testing.T) {
	assert.EqualBytes(a, b, newCallback(t))
}

// EqualError invariant will call t.Error() if two errors are not equal.
func EqualError(a, b error, t *testing.T) {
	assert.EqualError(a, b, newCallback(t))
}

// Bigger invariant will call t.Error() if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T, t *testing.T) {
	assert.Bigger(a, b, newCallback(t))
}

// Lesser invariant will call t.Error() if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T, t *testing.T) {
	assert.Lesser(a, b, newCallback(t))
}
