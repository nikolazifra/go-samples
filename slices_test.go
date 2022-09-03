package samples

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Equals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, val := range a {
		if !slices.Contains[string](b, val) {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	a := []string{"one"}
	b := []string{"two", "three"}
	want := []string{"one", "two", "three"}
	if got := Merge(a, b); !Equals(want, got) {
		t.Errorf("Hello() = want %q, got %q", want, got)
	}
}

func TestMergeDups(t *testing.T) {
	a := []string{"one", "two", "four"}
	b := []string{"two", "three", "one", "two"}
	want := []string{"one", "two", "four", "three"}
	if got := Merge(a, b); !Equals(want, got) {
		t.Errorf("Hello() = want %q, got %q", want, got)
	}
}
