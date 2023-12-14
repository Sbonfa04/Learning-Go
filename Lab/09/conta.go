/*
Scrivere un programma conta.go, dotato di:

- una funzione ricorsiva recursiveCount(el int, list []int) int che
restituisca il numero di volte che el appare tra i valori di list

- una funzione main() che legga da standard input (ctrl D per terminare)
una lista di numeri interi (che posono essere positivi, negativi, nulli)
ed emetta nel flusso d'uscita quante volte il primo numero letto è ripetuto
nella sequenza dei successivi numeri letti.
*/

package main

import (
	"fmt"
)

func recursiveCount(el int, list []int) int {
	if len(list) == 0 {
		return -1
	}

	if list[0] == el {
		return 1 + recursiveCount(el, list[1:])
	}

	return recursiveCount(el, list[1:])
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
	fmt.Println("il primo numero è apparso", recursiveCount(n, list), "volte")
}
