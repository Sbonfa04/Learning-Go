package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cancella(lista []string) []string {
	var lista2 []string

	for i := 0; i < len(lista); i++ {
		n, err := strconv.Atoi(lista[i])
		if err != nil {
			lista2 = append(lista2, lista[i])
		} else {
			if (i + n + 1) >= len(lista) {
				return lista2
			}
		}
	}
	return lista2

}

func main() {
	var lista []string

	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
	} else {
		_, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("File non accessibile")
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				parola := scanner.Text()
				lista = append(lista, parola)
			}

			fmt.Println(lista)
			fmt.Println(cancella(lista))
		}
	}
}
