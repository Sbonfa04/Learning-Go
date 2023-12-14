/*
Funzioni variadiche, ovvero funzioni che possono avere un numero
variabile di parametri.
*/
package main

import "fmt"

func somma1(x ...int) int { //con i puntini indico che la funzione può avere un numero variabile di parametri
	var s int
	for _, v := range x {
		s += v
	}
	return s
}

func somma2(x []int) int { //altro modo per fare la stessa cosa ma con un array
	var s int
	for _, v := range x {
		s += v
	}
	return s
}

func main() {

	x := somma1(5, 7, 8, 9, 4, 3, 2, 1, 5, 6, 7, 8, 9, 4, 3, 2, 1, 5, 6, 7, 8, 9, 4, 3, 2, 1, 5, 6, 7, 8, 9, 4, 3, 2, 1, 5, 6, 7, 8, 9, 4, 3, 2, 1)
	y := somma2([]int{5, 6, 7, 8, 9})

	fmt.Println(x)
	fmt.Println(y)

	var s, t, r []string
	s = append(s, "ciao", "come", "stai")
	t = append(t, "pippo", "pluto", "paperino")
	s = append(s, t...) //con i puntini indico che voglio aggiungere tutti gli elementi di t a in fondo ad s

	fmt.Println(s)

	r = append(s[:2], s[3:]...) //con i puntini indico che voglio eliminare l'elemento in posizione 2 di s

	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(r)
}
