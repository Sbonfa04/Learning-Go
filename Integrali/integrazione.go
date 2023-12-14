package main

import (
	"fmt"
	"math"
)

func integral(f func(float64) float64, a, b float64, n int) float64 {
	var sum float64 = 0
	var delta float64 = (b - a) / float64(n)
	for i := 0; i < n; i++ {
		x0 := a + delta*float64(i)
		x1 := a + delta*float64(i+1)
		y0 := f(x0)
		y1 := f(x1)
		area := (y0 + y1) * delta / 2
		sum += area
	}
	return sum
}

func parabola(x float64) float64 {
	return x * x
}

func main() {
	a := integral(parabola, 0, 1, 1000)
	b := integral(math.Sin, 0, math.Pi, 1000)
	fmt.Println(a, b)
}
