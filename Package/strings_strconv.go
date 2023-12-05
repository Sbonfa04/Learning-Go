package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	//"math"
	_ "bufio"
	_ "os"
)

func main() {

	const VOCALI = "aeiouAEIOU"
	const s1 = "c"
	const s2 = "u"
	var parola string

	fmt.Scan(&parola)

	if strings.Contains(parola, s2) {
		fmt.Printf("(1) la parola %q contiene %q\n", parola, s2)
	} else {
		fmt.Printf("(1) la parola %q NON contiene %q\n", parola, s2)
	}

	if strings.ContainsAny(parola, VOCALI) {
		fmt.Printf("(2) la parola %q continene almeno una vocale\n", parola)
	} else {
		fmt.Printf("(2) la parola %q NON continene una vocale\n", parola)
	}

	x := strings.Count(parola, s1)
	fmt.Printf("(3) la parola %q contiene %v volte %q\n", parola, x, s1)

	primaVocale := -1
	ultimaVocale := -1

	for i, c := range parola {
		if strings.ContainsRune(VOCALI, c) {
			if primaVocale == -1 {
				primaVocale = i + 1
			}
			ultimaVocale = i + 1
		}
	}

	fmt.Printf("(4) La prima vocale si trova alla posizione %d; ", primaVocale)
	fmt.Printf("L'ultima vocale si trova alla posizione %d\n", ultimaVocale)

	ripetuto := strings.Repeat(s2, 3)
	fmt.Println("(5)", ripetuto)

	ripetuto2 := strings.Repeat(s1, 5)
	fmt.Println("(6)", ripetuto2)

	cifreParola := ""
	for _, c := range parola {
		if unicode.IsDigit(c) {
			cifreParola += string(c)
		}
	}

	fmt.Printf("(7) cifre estratte da %v: %s\n", parola, cifreParola)

	n, _ := strconv.ParseInt(cifreParola, 10, 64)
	fmt.Printf("(8) intero %d\n", n)

	f, _ := strconv.ParseFloat(cifreParola, 64)
	fmt.Printf("(9) float %f\n", f)
}
