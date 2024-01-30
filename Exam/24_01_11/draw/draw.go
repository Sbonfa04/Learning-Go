/*
draw
------

Scrivere un programma 'draw.go' dotato di:

- una funzione
    DrawPoint(c byte, k int) string

che restituisce una stringa formata da k spazi bianchi
seguiti dal carattere c

- una funzione
    DrawSegment(c byte, k, l int) string

che restituisce una stringa formata da k spazi bianchi
seguiti da l caratteri c

- una funzione
    main()

che legge come parametri da linea di comando (in quest'ordine)
due numeri interi l e n e un carattere (byte) c,
e, se l e n sono > 0, produce su standard output una scala di n gradini
di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
altrimenti non fa niente.

Si può assumere che il programma venga lanciato con tre parametri
dei tipi attesi (non occorre cioè fare controlli).

Esempi
------

$ ./draw 3 1 x
xxx

	x
	x

$ ./draw 3 2 +
+++

	+
	+
	+++
	  +
	  +
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func DrawPoint(c byte, k int) string {
	var s string
	for i := 0; i < k; i++ {
		s += " "
	}
	s += string(c)
	return s
}

func DrawSegment(c byte, k, l int) string {
	var s string
	for i := 0; i < k; i++ {
		s += " "
	}
	for i := 0; i < l; i++ {
		s += string(c)
	}
	return s
}

func main() {
	l, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	c := os.Args[3][0]

	if l > 0 && n > 0 {
		for i := 0; i < n; i++ {
			fmt.Println(DrawSegment(c, i*(l-1), l))
			for j := 1; j < l; j++ {
				fmt.Println(DrawPoint(c, l+i*(l-1)-1))
			}
		}
	}
}
