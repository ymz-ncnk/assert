package asserterror

import (
	"fmt"
	"testing"
	"time"

	"github.com/ymz-ncnk/assert"
	"golang.org/x/exp/constraints"
)

const msgSep = ": "

func newCallback(t *testing.T, msgAndArgs ...any) assert.Callback {
	return func(msg string) {
		t.Helper()
		if len(msgAndArgs) > 0 {
			if len(msgAndArgs) == 1 {
				msg += msgSep + fmt.Sprint(msgAndArgs[0])
			} else {
				msg += msgSep + fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
			}
		}
		t.Error(msg)
	}
}

// Equal invariant will call t.Error() if two compared variables are not equal.
func Equal[T comparable](t *testing.T, a, b T, msgAndArgs ...any) {
	t.Helper()
	assert.Equal(a, b, newCallback(t, msgAndArgs...), t)
}

// EqualDeep invariant will call t.Error() if two variables are not deep equal.
func EqualDeep(t *testing.T, a, b any, msgAndArgs ...any) {
	t.Helper()
	assert.EqualDeep(a, b, newCallback(t, msgAndArgs...), t)
}

// EqualBytes invariant will call t.Error() if two byte slices are not equal.
func EqualBytes(t *testing.T, a, b []byte, msgAndArgs ...any) {
	t.Helper()
	assert.EqualBytes(a, b, newCallback(t, msgAndArgs...), t)
}

// EqualError invariant will call t.Error() if two errors are not equal.
func EqualError(t *testing.T, a, b error, msgAndArgs ...any) {
	t.Helper()
	assert.EqualError(a, b, newCallback(t, msgAndArgs...), t)
}

// Bigger invariant will call t.Error() if a is lesser than b.
func Bigger[T constraints.Ordered](t *testing.T, a, b T, msgAndArgs ...any) {
	t.Helper()
	assert.Bigger(a, b, newCallback(t, msgAndArgs...), t)
}

// Lesser invariant will call t.Error() if a is bigger than b.
func Lesser[T constraints.Ordered](t *testing.T, a, b T, msgAndArgs ...any) {
	t.Helper()
	assert.Lesser(a, b, newCallback(t, msgAndArgs...), t)
}

// SameTime invariant will panic if two times are not equal.
func SameTime(t *testing.T, a, b time.Time, delta time.Duration,
	msgAndArgs ...interface{}) {
	t.Helper()
	assert.SameTime(a, b, delta, newCallback(t, msgAndArgs...), nil)
}
