/*
scrivere una funzione ricorsiva che data una stringa
stabilisce se è una palindroma
*/
package main

import (
	"fmt"
)

func isPalindroma(s string) bool {
	if len(s) <= 1 {
		return true
	} else {
		return s[0] == s[len(s)-1] && isPalindroma(s[1:len(s)-1])
	}
}

func main() {
	var x string
	fmt.Scan(&x)
	if isPalindroma(x) == true {
		fmt.Printf("la parola %v è palindroma", x)
	} else {
		fmt.Printf("la stringa %v non è palindorma", x)
	}
}
