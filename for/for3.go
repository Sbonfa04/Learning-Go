//stampa i multipli di 3 minori di x (input)
package main

import (
	"fmt"
)

func main() {

	var x int
	fmt.Scan(&x)
	for i := 3; i < x; i += 3 {
		fmt.Println(i)
	}
}
