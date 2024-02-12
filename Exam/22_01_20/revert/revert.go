/*

Revert
------

Realizzare un programma Go (nome file 'revert.go') che legge da standard input
**stringhe** costituite da soli caratteri **ASCII standard (a 7 bit)**
(separate da un numero arbitrario di white space) e stampa solo quelle di lunghezza pari,
al contrario (il primo carattere per ultimo e l'ultimo per primo, ecc.), una per riga.

Il programma termina quando legge la stringa "stop", che non va trattata.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "stop" {
			break
		}
		if len(s)%2 == 0 {
			for i := len(s) - 1; i >= 0; i-- {
				fmt.Printf("%c", s[i])
			}
			fmt.Println()
		}
	}
}
