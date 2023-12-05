package main

import (
	"fmt"
)

func main() {
	var p Punto
	p = NewPunto(3.14, 2.7)
	fmt.Println(Str(p))
}
