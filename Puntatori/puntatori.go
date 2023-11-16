package main

import (
	"fmt"
)

func main() {

	var x int
	var p *int
	var q **int

	x = 7
	p = &x
	q = &p
	*p = 50            //x = 50
	**q = *p - (x + 1) //x = 50 - (50+1)
	*p++               //x + 1
	**q--              //x - 1
	fmt.Println(x)     //x = -1
}
