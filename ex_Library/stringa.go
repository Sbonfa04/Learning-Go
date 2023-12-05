package main

import (
	"fmt"
)

func Str(p Punto) string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}
