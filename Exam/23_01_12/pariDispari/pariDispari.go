/*
Scrivere un programma 'pariDispari.go' che legge una sequenza di valori interi da linea di
comando e controlla che si alternino valori pari e valori dispari.
In questo caso il programma stampa il messaggio "sequenza valida",
altrimenti "elemento in posizione i non valido",
dove i è la posizione del primo  elemento (da sinistra) che non rispetta la regola
di alternanza o che non è un valore numerico.

In caso di mancanza di valori, il programma deve stampare "nessun valore in ingresso".

La sequenza può iniziare sia con un valore pari sia con uno dispari.

Si ricorda che lo zero è un numero pari.

Esempi
------

Argomenti:

	3 8 1 12

Output:

	sequenza valida

Argomenti:

	4

Output:

	sequenza valida

Argomenti:

	1 2 3 5

Output:

	elemento in posizione 4 non valido

Argomenti:

	1 2 3eqeqw 5

Output:

	elemento in posizione 3 non valido

Argomenti:

Output:

	nessun valore in ingresso
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("nessun valore in ingresso")
		return
	}
	nums := make([]int, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("elemento in posizione %d non valido\n", i+1)
			return
		}
		nums[i] = num
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			fmt.Printf("elemento in posizione %d non valido\n", i+2)
			return
		}
	}
	fmt.Println("sequenza valida")
}
