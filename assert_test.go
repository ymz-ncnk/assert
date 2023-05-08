package assert

import (
	"errors"
	"fmt"
	"testing"
)

type serr struct {
	s string
}

func (e *serr) Error() string {
	return e.s
}

func TestAssert(t *testing.T) {
	On = true

	t.Run("Equal", func(t *testing.T) {
		Equal(1, 1)
		defer func() {
			if str := recover(); str != "1 != 2" {
				fmt.Println(str)
				t.Errorf("unexpected panic")
			}
		}()
		Equal(1, 2)
	})

	t.Run("EqualError", func(t *testing.T) {
		var (
			err1 = errors.New("some")
			err2 = errors.New("some")
		)
		EqualError(err1, err2)
		EqualError(err1, err1)
		EqualError(nil, nil)

	})

	t.Run("EqualError different type", func(t *testing.T) {
		var (
			err1 = errors.New("some")
			err3 = &serr{"some"}
		)
		defer func() {
			if str := recover(); str != "*errors.errorString != *assert.serr" {
				t.Errorf("unexpected panic")
			}
		}()
		EqualError(err1, err3)
	})

	t.Run("EqualError different msg", func(t *testing.T) {
		var (
			err1 = errors.New("some")
			err4 = errors.New("another")
		)
		defer func() {
			if str := recover(); str != "some != another" {
				t.Fatal("unexpected panic")
			}
		}()
		EqualError(err1, err4)
	})

	t.Run("EqualDeep", func(t *testing.T) {
		var (
			sl1 = []string{"1"}
			sl2 = []string{"1"}
			sl3 = []string{"2"}
		)
		EqualDeep(sl1, sl2)
		defer func() {
			if str := recover(); str != "[1] != [2]" {
				t.Errorf("unexpected panic")
			}
		}()
		EqualDeep(sl1, sl3)
	})

	t.Run("Bigger", func(t *testing.T) {
		Bigger(2, 1)
		defer func() {
			if str := recover(); str != "1 < 2" {
				fmt.Println(str)
				t.Errorf("unexpected panic")
			}
		}()
		Bigger(1, 2)
	})

	t.Run("Lesser", func(t *testing.T) {
		Lesser(1, 2)
		defer func() {
			if str := recover(); str != "2 > 1" {
				fmt.Println(str)
				t.Errorf("unexpected panic")
			}
		}()
		Lesser(2, 1)
	})
}
