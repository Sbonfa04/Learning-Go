package main

import (
	"fmt"
)

func main() {
	var x int
	var p *int
	var q **int
	x = 7
	p = &x // punta x
	q = &p // punta x
	p = new(int) // si riferisce ad una nuova memoria
	*p = 50 // nuova memoria = 50
	q = new(*int) // atra memoria
	*q = &x // altra memoria punta x
	**q = 5 // x = 5
	*q = new(int) // altra memoria nuova associzione
	**q = 12 // altra memoria nuova associazione = 12
	fmt.Println(x, *p, **q) //5, 50, 12
}
