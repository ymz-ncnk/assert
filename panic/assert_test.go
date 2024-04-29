package assert

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
)

func TestEqualOk(t *testing.T) {
	Equal(1, 1)
}

func TestEqualFail(t *testing.T) {
	defer func() {
		if str := recover(); str != "1 != 2" {
			fmt.Println(str)
			t.Errorf("unexpected panic")
		}
	}()
	Equal(1, 2)
}

func TestEqualErrorOk(t *testing.T) {
	EqualError(errors.New("some"), errors.New("some"))
}

func TestEqualErrorFailDifferentType(t *testing.T) {
	defer func() {
		if str := recover(); str != "*errors.errorString != *url.Error" {
			t.Errorf("unexpected panic")
		}
	}()
	EqualError(errors.New("some"), &url.Error{})
}

func TestEqualErrorFailDifferentStr(t *testing.T) {
	defer func() {
		if str := recover(); str != "some != another" {
			t.Errorf("unexpected panic")
		}
	}()
	EqualError(errors.New("some"), errors.New("another"))
}

func TestEqualDeepOk(t *testing.T) {
	EqualDeep([]int{1}, []int{1})
}

func TestEqualDeepFail(t *testing.T) {
	defer func() {
		if str := recover(); str != "[1] != [2]" {
			t.Errorf("unexpected panic")
		}
	}()
	EqualDeep([]int{1}, []int{2})
}

func TestBiggerOk(t *testing.T) {
	Bigger(2, 1)
}

func TestBiggerFail(t *testing.T) {
	defer func() {
		if str := recover(); str != "1 < 2" {
			fmt.Println(str)
			t.Errorf("unexpected panic")
		}
	}()
	Bigger(1, 2)
}

func TestLesserOk(t *testing.T) {
	Lesser(1, 2)
}

func TestLesserFail(t *testing.T) {
	defer func() {
		if str := recover(); str != "2 > 1" {
			fmt.Println(str)
			t.Errorf("unexpected panic")
		}
	}()
	Lesser(2, 1)
}
