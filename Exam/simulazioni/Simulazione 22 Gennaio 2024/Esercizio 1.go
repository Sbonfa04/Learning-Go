package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prodottoVettoriale() {

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a := make([]int, len(strings.Fields(scanner.Text())), len(strings.Fields(scanner.Text())))

	for i, v := range strings.Fields(scanner.Text()) {
		value, _ := strconv.Atoi(v)
		a[i] = value
	}

	scanner.Scan()
	b := make([]int, len(strings.Fields(scanner.Text())), len(strings.Fields(scanner.Text())))

	for i, v := range strings.Fields(scanner.Text()) {
		value, _ := strconv.Atoi(v)
		b[i] = value
	}

	matrice := make([][]int, len(a))

	for i := 0; i < len(a); i++ {
		matrice[i] = make([]int, len(b))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			matrice[i][j] = a[i] * b[j]
		}
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(matrice[i])
	}

}
