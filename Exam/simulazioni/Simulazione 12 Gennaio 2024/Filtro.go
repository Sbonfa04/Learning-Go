package main

import (
	"fmt"
	"os"
)

func main() {
	numero := os.Args[1]

	for i := 0; i < len(numero)-1; i++ {
		if numero[i] < numero[i+1] {
			fmt.Print(string(numero[i])) //rune to string
		}
	}
}
