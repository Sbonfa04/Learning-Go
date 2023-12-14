package main

import (
	"testing"
)

func TestMedian(t *testing.T) {
	var p1, p2, pm Point
	p1 = NewPoint(0, 0)
	p2 = NewPoint(2, 4)
	pm = p1.Median(p1, p2)

	if pm.X != 1 || pm.Y != 2 {
		t.Error("Expected 1,2 got ", pm.X, pm.Y)
	}
}
