/*
Scrivere un programma che riceve una stringa da riga di comando e stampa la stringa di N
caratteri in un quadrato di dimensione NXN. Si assuma che la stringa utilizzi solo caratteri ASCII.
Esempio d’esecuzione:
$ go run esercizio_3.go mano
m  m

	aa
	nn

o  o
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	parola := os.Args[1]
	n := len(parola)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print(string(parola[j]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
