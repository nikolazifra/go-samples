package samples

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Equals[K comparable](a, b []K) bool {
	if len(a) != len(b) {
		return false
	}

	for _, val := range a {
		if !slices.Contains(b, val) {
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
		t.Errorf("Merge() = want %q, got %q", want, got)
	}
}

func TestMergeInt(t *testing.T) {
	a := []int{1}
	b := []int{2, 3}
	want := []int{1, 2, 3}
	if got := Merge(a, b); !Equals(want, got) {
		t.Errorf("Merge() = want %q, got %q", want, got)
	}
}

func TestMergeFloat(t *testing.T) {
	a := []float64{1.0}
	b := []float64{2.0, 3.00}
	want := []float64{1, 2, 3}
	if got := Merge(a, b); !Equals(want, got) {
		t.Errorf("Merge() = want %v, got %v", want, got)
	}
}

func TestMergeDups(t *testing.T) {
	a := []string{"one", "two", "four"}
	b := []string{"two", "three", "one", "two"}
	want := []string{"one", "two", "four", "three"}
	if got := Merge(a, b); !Equals(want, got) {
		t.Errorf("Merge() = want %q, got %q", want, got)
	}
}
