/*
Scrivere un programma righello.go, dotato di:

- una funzione ricorsiva righello(n int) che, data la dimensione di un righello,
emetta nel flusso d'uscita un righello della dimensione data
- una funzione main() che, data la dimensione come argomento sulla linea di comando,
produca il righello corrispondente

La dimensione può essere 0, nel qual caso il programma non emette nulla.
Per la definizione di "righello" si legga con attenzione quanto segue:

chiamiamo "tacca" di lunghezza N una sequenza di N caratteri "-" seguiti da "a-capo"
chiamiamo "righello" di dimensione M > 0 una tacca lunga M preceduta e seguita da un
righello di dimensione M - 1, assumendo per convenzione che un righello di dimensione 0
non contenga alcuna tacca (o altro righello)
*/

package main

import "fmt"

func righello(n int) {
	if n <= 0 {
		return
	}
	if n == 1 {
		fmt.Println("-")
	} else {
		righello(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("-")
		}
		fmt.Println()
		righello(n - 1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	righello(n)
}
