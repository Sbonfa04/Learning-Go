/*
Scrivere un programma array.go dotato di:

- una costante DIM = 5 per la dimensione dell'array, dichiarata a livello di package

- una funzione main che inizializza a piacere un array di int di dimensione DIM (ad esempio con numeri 0, 1, 2, ...) e testa le due seguenti funzioni che lavorano **direttamente sull'array stesso**, senza quindi dichiarare e usare altri array. Il programma stampa l'array iniziale, l'array dopo aver invocato reverse e l'array dopo aver invocato switchFirstLast.

- una funzione reverse che inverte l'ordine dei valori in un array di dimensione DIM, mettendo il primo in fondo, il secondo in penultima posizione e così via, nell'array stesso

- una funzione switchFirstLast che scambia il primo con l'ultimo dei valori in un array di dimensione DIM nell'array stesso


Esempio di esecuzione

$ go run array.go

array: [0 1 2 3 4]
reverse: [4 3 2 1 0]
switchFirstLast: [0 3 2 1 4]

*/

package main

import (
	"fmt"
)

const DIM = 5

func reverse(array [DIM]int) {
	for i, j := 0, DIM-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return
}

func switchFirstLast(array [DIM]int) {
	array[0], array[DIM-1] = array[DIM-1], array[0]
	return
}

func main() {

	array := [DIM]int{10, 20, 30, 40, 50}

	// Stampa l'array iniziale
	fmt.Printf("array: %v\n", array)

	// Inverte l'ordine degli elementi nell'array
	reverse(&array)
	fmt.Printf("reverse: %v\n", array)

	// Ripristina l'array all'originale
	reverse(&array)

	// Scambia il primo e l'ultimo elemento nell'array
	switchFirstLast(&array)
	fmt.Printf("switchFirstLast: %v\n", array)
}
