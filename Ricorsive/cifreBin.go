/*dato un numero in decimale calcola la quantità di cifre in binario*/
package main

import (
	"fmt"
)

func f(x int) int {
	if x < 2 {
		return 1
	} else {
		return 1 + f(x/2)
	}
}

// dato un numero lo trasforma in binario (dritto)
func bin(x int) {
	if x < 2 {
		fmt.Println(x)
	} else {
		resto := x % 2
		bin(x / 2)
		fmt.Println(resto)
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	x = f(x)
	fmt.Println(x)
}
