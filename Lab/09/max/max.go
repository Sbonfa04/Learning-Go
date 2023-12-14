/*
Scrivere un programma max.go, dotato di:

- una funzione ricorsiva recursiveMax(list []int) int che restituisca il massimo
tra i valori di list

- una funzione main() che legga da standard input (ctrl D per terminare) una lista
di numeri interi (che posono essere positivi, negativi, nulli) ed emetta nel flusso
d'uscita il massimo tra i numeri letti.

Il massimo di una sequenza vuota non è definito, quindi assumiamo (non è richiesto
che vengano fatti controlli) che la sequenza abbia sempre almeno un numero.
*/

package main

import (
	"fmt"
)

func recursiveMax(list []int) int {
	var max int
	if len(list) == 1 {
		return list[0]
	} else {
		max = recursiveMax(list[1:])
		if list[0] > max {
			return list[0]
		} else {
			return max
		}
	}
}

func main() {
	var n int
	var list []int

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		list = append(list, n)
	}
	fmt.Println(recursiveMax(list))
}
