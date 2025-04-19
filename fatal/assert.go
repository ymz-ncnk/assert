package assertfatal

import (
	"testing"

	"github.com/ymz-ncnk/assert"
	"golang.org/x/exp/constraints"
)

func newCallback(t *testing.T) assert.Callback {
	return func(msg string) { t.Helper(); t.Fatal(msg) }
}

// Equal invariant will call t.Fatal() if two compared variables are not equal.
func Equal[T comparable](a, b T, t *testing.T) {
	t.Helper()
	assert.Equal(a, b, newCallback(t), t)
}

// EqualDeep invariant will call t.Fatal() if two variables are not deep equal.
func EqualDeep(a, b interface{}, t *testing.T) {
	t.Helper()
	assert.EqualDeep(a, b, newCallback(t), t)
}

// EqualBytes invariant will call t.Fatal() if two byte slices are not equal.
func EqualBytes(a, b []byte, t *testing.T) {
	t.Helper()
	assert.EqualBytes(a, b, newCallback(t), t)
}

// EqualError invariant will call t.Fatal() if two errors are not equal.
func EqualError(a, b error, t *testing.T) {
	t.Helper()
	assert.EqualError(a, b, newCallback(t), t)
}

// Bigger invariant will call t.Fatal() if a is lesser than b.
func Bigger[T constraints.Ordered](a, b T, t *testing.T) {
	t.Helper()
	assert.Bigger(a, b, newCallback(t), t)
}

// Lesser invariant will call t.Fatal() if a is bigger than b.
func Lesser[T constraints.Ordered](a, b T, t *testing.T) {
	t.Helper()
	assert.Lesser(a, b, newCallback(t), t)
}
