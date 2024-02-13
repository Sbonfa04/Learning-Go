package main

import (
	"fmt"
	"os"
)

func main() {
	parola := os.Args[1]
	var lista []string
	for c, i := range parola {
		lista[i] = string(c)
	}

	for i := 0; i < len(lista); i++ {
		for j := 0; j < len(lista); j++ {
			if j == len(lista) {
				j = 0
			}
			fmt.Print(lista[j])
		}
	}
}
