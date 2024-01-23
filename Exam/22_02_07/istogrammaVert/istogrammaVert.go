/*
ISTOGRAMMA
----------

Scrivere un programma (il file deve chiamarsi 'istogrammaVert.go') che legge da linea di comando una
serie di numeri interi non negativi e ne disegna il relativo istogramma. L'istogramma sarà formato da
colonne di asterischi che rappresentano i valori letti.
Non sono richiesti controlli sui dati.

Esempio:

$ ./istogrammaVert 5 0 2 9 4
   *
   *
   *
   *
*  *
*  **
*  **
* ***
* ***

$ ./istogrammaVert 2 3 4 5
   *
  **
 ***
****
****


*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := make([]int, len(os.Args)-1)
	max := 0
	for i, arg := range os.Args[1:] {
		num, _ := strconv.Atoi(arg)
		nums[i] = num
		if num > max {
			max = num
		}
	}

	for i := max; i > 0; i-- {
		for _, num := range nums {
			if num >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
