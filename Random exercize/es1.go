package main

import (
	. "fmt"
)

func main() {

	var n, x1, x2, curr int
	Scan(&n)

	x1 = 1
	x2 = 1

	for i := 0; i < n; i++ {
		if i < 2 {
			curr = 1
		} else {
			curr = x1 + x2
		}
		for j := 0; j < curr; j++ {
			Print("*")
		}
		x1, x2 = x2, curr
	}

}
