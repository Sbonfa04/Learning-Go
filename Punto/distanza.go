package main

import (
	"math"
)

func Dist(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}
