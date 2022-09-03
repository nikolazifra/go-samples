package samples

import (
	"testing"
)

func TestRoots(t *testing.T) {
	var a float64 = 2
	var b float64 = 10
	var c float64 = 8
	var x1 float64 = -1
	var x2 float64 = -4
	if y1, y2 := findRoots(a, b, c); x1 != y1 || x2 != y2 {
		t.Errorf("findRoots() = wanted %f,%f but got %f,%f", x1, x2, y1, y2)
	}
}

func TestRootsOpt(t *testing.T) {
	var a float64 = 2
	var b float64 = 10
	var c float64 = 8
	var x1 float64 = -1
	var x2 float64 = -4
	if y1, y2 := findRootsOpt(a, b, c); x1 != y1 || x2 != y2 {
		t.Errorf("findRoots() = wanted %f,%f but got %f,%f", x1, x2, y1, y2)
	}
}
