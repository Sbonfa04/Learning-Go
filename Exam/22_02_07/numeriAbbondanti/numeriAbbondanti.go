/*
NUMERI ABBONDANTI
-----------------

Scrivere un programma (il file deve chiamarsi 'numeriAbbondanti.go') che legge un numero intero
positivo n da standard input e stampa i primi n "numeri abbondanti", uno per riga.

Un numero abbondante è un numero naturale minore della somma dei suoi divisori interi
(1 compreso, numero stesso ovviamente escluso).

Per esempio, 12 è un numero abbondante poiché inferiore alla somma dei suoi divisori: 1+2+3+4+6=16
mentre invece 15 non lo è in quanto 1+3+5 = 9

Il programma deve essere dotato di una funzione
	Abbondante(n int) bool
che, dato un intero positivo, restituisce true se n è un numero abbondante, false altrimenti.
Se il numero passato come parametro è minore o uguale a 0, la funzione restituisce false.

Esempio di esecuzione
---------------------

./numeriAbbondanti
3
12
18
20

NOTA BENE: il 3 alla prima riga è l'input, i numeri successivi sono l'output

*/

package main

import (
	"fmt"
)

func Abbondante(n int) bool {
	if n <= 0 {
		return false
	}
	var somma int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			somma += i
		}
	}
	return somma > n
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; n > 0; i++ {
		if Abbondante(i) {
			fmt.Println(i)
			n--
		}
	}
}
