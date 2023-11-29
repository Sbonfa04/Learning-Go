package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	x = fatt(x)
	fmt.Println(x)
}

func fatt(n int) (ris int) {
	if n == 0 {
		ris = 1
	} else {
		s := fatt(n - 1) //n-1 diventa l'input della funzione ricorsiva fatt
		ris = n * s      //e si entra in una specie di loop fino a n = 0
	}
	return
	/* al posto che creare una variabile "ris" posso anche solo fare un return con
	   il valore di "ris" nei due casi dell'if*/
}
