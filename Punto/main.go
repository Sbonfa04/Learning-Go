package main

import (
	"fmt"
)

func main() {
	var p1, p2 Punto
	p1 = NewPunto(1.3, 2)
	p2 = NewPunto(3.1, 4)
	fmt.Println(Dist(p1, p2))
	fmt.Println(Media(p1, p2))
	fmt.Println(Str(p1), Str(p2))
}
