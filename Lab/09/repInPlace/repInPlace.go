package main

import (
	"fmt"
	"os"
	"strings"
)

func repInPlace(stringa []rune, old rune, new rune) {
	for i, c := range stringa {
		if c == old {
			stringa[i] = new
		}
	}
}

func main() {
	var modString string
	var oldChar rune
	var newChar rune

	lunghezza := len(os.Args)
	if lunghezza < 4 {
		return
	} else {
		modString = strings.Join(os.Args[1:lunghezza-2], " ")
		oldChar = []rune(os.Args[lunghezza-2])[0]
		newChar = []rune(os.Args[lunghezza-1])[0]
		fmt.Println(modString)
		modStringRune := []rune(modString)
		repInPlace(modStringRune, oldChar, newChar)
		fmt.Println(string(modStringRune))
	}
}
