package assert

import (
	"errors"
	"net/url"
	"testing"
)

func TestEqualOk(t *testing.T) {
	if !Equal(1, 1, func(msg string) { t.Error("unexpected result") }, nil) {
		t.Error("unexpected result")
	}
}

func TestEqualFail(t *testing.T) {
	if Equal(1, 2, func(msg string) {}, nil) {
		t.Error("unexpected result")
	}
}

func TestEqualErrorOk(t *testing.T) {
	if !EqualError(errors.New("some"), errors.New("some"), func(msg string) {
		t.Error("unexpected result")
	}, nil) {
		t.Error("unexpected result")
	}
}

func TestEqualErrorFailDifferentType(t *testing.T) {
	var (
		count = 0
		ptr   = &count
	)
	if EqualError(errors.New("some"), &url.Error{}, func(msg string) {
		*ptr += 1
	}, nil) || count != 1 {
		t.Error("unexpected result")
	}
}

func TestEqualErrorFailDifferentStr(t *testing.T) {
	var (
		count = 0
		ptr   = &count
	)
	if EqualError(errors.New("some"), errors.New("another"), func(msg string) {
		*ptr += 1
	}, nil) || count != 1 {
		t.Error("unexpected result")
	}
}

func TestEqualDeepOk(t *testing.T) {
	var (
		s1 = []string{"1"}
		s2 = []string{"1"}
	)
	if !EqualDeep(s1, s2, func(msg string) { t.Error("unexpected result") }, nil) {
		t.Error("unexpected result")
	}
}

func TestEqualDeepFail(t *testing.T) {
	var (
		s1 = []string{"1"}
		s2 = []string{"2"}
	)
	if EqualDeep(s1, s2, func(msg string) {}, nil) {
		t.Error("unexpected result")
	}
}

func TestBiggerOk(t *testing.T) {
	if !Bigger(2, 1, func(msg string) { t.Error("unexpected result") }, nil) {
		t.Error("unexpected result")
	}
}

func TestBiggerFail(t *testing.T) {
	if Bigger(1, 2, func(msg string) {}, nil) {
		t.Error("unexpected result")
	}
}

func TestLesserOk(t *testing.T) {
	if !Lesser(1, 2, func(msg string) { t.Error("unexpected result") }, nil) {
		t.Error("unexpected result")
	}
}

func TestLesserFail(t *testing.T) {
	if Lesser(2, 1, func(msg string) {}, nil) {
		t.Error("unexpected result")
	}
}
