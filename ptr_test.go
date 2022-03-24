package ptr_test

import (
	"ptr"
	"reflect"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	if expected, got := reflect.TypeOf(new(int)), reflect.TypeOf(ptr.Convert(0)); expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
	if expected, got := reflect.TypeOf(new(string)), reflect.TypeOf(ptr.Convert("")); expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
	if expected, got := reflect.TypeOf(new(time.Time)), reflect.TypeOf(ptr.Convert(time.Time{})); expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestString(t *testing.T) {
	if expected, got := "0", ptr.String(new(int)); expected != got {
		t.Errorf("expected: '%v', got: '%v'", expected, got)
	}
	if expected, got := "nil", ptr.String[*int](nil); expected != got {
		t.Errorf("expected: '%v', got: '%v'", expected, got)
	}
	if expected, got := "", ptr.String(new(string)); expected != got {
		t.Errorf("expected: '%v', got: '%v'", expected, got)
	}
	if expected, got := "nil", ptr.String[*string](nil); expected != got {
		t.Errorf("expected: '%v', got: '%v'", expected, got)
	}
}
