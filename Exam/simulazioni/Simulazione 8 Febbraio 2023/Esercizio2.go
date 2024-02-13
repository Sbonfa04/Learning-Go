/*
CXLI
CX C>X? SI, somma +=100
XL X > L? NO, somma += L-x = 50-10 = 140 Avanti due
I somma += 1 = 141


MCMLXXXIII
MC M>C? SI, somma += M = 1000
CM C>M? NO, somma += M-C = 1000-100 = 1900 Avanti due
LX L>X? SI, somma += L = 1950
XX X>X? NO, MA X=X somma += X+X = 1970 avanti due
XI X>I? SI, somma += X = 1980
II I>I? NO, MA I=I somma+= I+I = 1982 Avanti due
I somma += I = 1983

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	mappaturaRomana := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'M': 1000, 'C': 100, 'D': 500}

	stringaRomana := os.Args[1]
	somma := 0

	for i := 0; i < len(stringaRomana); i++ {
		if i+1 < len(stringaRomana) {
			if mappaturaRomana[stringaRomana[i]] > mappaturaRomana[stringaRomana[i+1]] {
				somma += mappaturaRomana[stringaRomana[i]]
			} else if mappaturaRomana[stringaRomana[i]] < mappaturaRomana[stringaRomana[i+1]] {
				somma += mappaturaRomana[stringaRomana[i+1]] - mappaturaRomana[stringaRomana[i]]
				if i != 0 {
					i += 1
				}

			} else {
				somma += mappaturaRomana[stringaRomana[i]] + mappaturaRomana[stringaRomana[i+1]]
				i += 1
			}
		} else {
			somma += mappaturaRomana[stringaRomana[i]]
		}
	}

	fmt.Println(somma)
}
