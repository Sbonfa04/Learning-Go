/*Scrivere una programma quadrati.go che legge da linea di comando una sequenza di
numeri interi non negativi e stampa solo quelli che sono dei quadrati (1, 4, 9, ecc.).
Il programma deve essere dotato di una funzione isSquare(n int) bool
che restituisce true se n è un quadrato, false altrimenti.*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func isSquare(n int) bool {

	if n == int(math.Sqrt(float64(n)))*int(math.Sqrt(float64(n))) {
		return true
	} else {
		return false
	}
}

func main() {

	for i := 1; i < (len(os.Args)); i++ {
		x, _ := strconv.Atoi(os.Args[i])
		if isSquare(x) {
			fmt.Println(x)
		}
	}
}
