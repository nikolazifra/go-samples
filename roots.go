package samples

import (
	"math"
)

func findRoots(a, b, c float64) (x1 float64, x2 float64) {
	sqr := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	x1 = (-b + sqr) / (2 * a)
	x2 = (-b - sqr) / (2 * a)
	return
}

func sq(a, b, c float64, ch chan float64) {
	x := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	ch <- x
	ch <- x
}

func x1(a, b float64, ch chan float64, r chan float64) {
	sqr := <-ch
	r <- ((-b + sqr) / (2 * a))
}

func x2(a, b float64, ch chan float64, r chan float64) {
	sqr := <-ch
	r <- ((-b - sqr) / (2 * a))
}

func findRootsOpt(a, b, c float64) (float64, float64) {
	ch := make(chan float64)
	rs := make(chan float64)
	go sq(a, b, c, ch)
	go x1(a, b, ch, rs)
	go x2(a, b, ch, rs)
	x := <-rs
	y := <-rs
	if x >= y {
		return x, y
	}
	return y, x
}
