/*
Scrivere un programma che riceve una stringa da riga di comando e stampa la stringa di N
caratteri in una X dove la parola parte dai due lati alti della X. Si assuma che la stringa utilizzi solo caratteri ASCII.
Esempio d’esecuzione:
$ go run filtro2.go ragno
r   r
 a a
  g
 n n
o   o
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	parola := os.Args[1]
	n := len(parola)
	for i := n - 1; i >= 0; i-- {
		fmt.Print(string(parola[i]))
	}
	fmt.Println()
}
