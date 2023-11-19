package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Print("inserisci un numero: ")
	x = math.MaxInt64

	for x >= 0 {
		fmt.Println(x)
		x--
	}
}
