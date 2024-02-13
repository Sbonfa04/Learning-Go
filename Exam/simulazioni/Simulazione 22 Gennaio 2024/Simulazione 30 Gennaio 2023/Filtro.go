package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	stringa := os.Args[1]
	righeDaStampare := 0
	spaziInMezzo := len(stringa)
	i := 0

	if len(stringa)%2 == 0 {
		righeDaStampare = len(stringa) / 2
	} else {
		righeDaStampare = (len(stringa) - 1) / 2

	}

	//stampo parte sopra
	for ; i < righeDaStampare; i++ {
		spaziInMezzo -= 2

		//spazi prima
		fmt.Print(strings.Repeat(" ", i))
		//carattere
		fmt.Print(string(stringa[i]))
		//spazi in mezzo
		fmt.Print(strings.Repeat(" ", spaziInMezzo))
		//carattere
		fmt.Print(string(stringa[i]))
		//a capo
		fmt.Println()
	}

	//se len è dispari, stampo carattere singolo
	if len(stringa)%2 != 0 {
		fmt.Print(strings.Repeat(" ", i))
		fmt.Print(string(stringa[i]))
		fmt.Println()
	}

	i -= 1
	//stampo parte sotto
	for ; i >= 0; i-- {

		//spazi prima
		fmt.Print(strings.Repeat(" ", i))
		//carattere
		fmt.Print(string(stringa[len(stringa)-i-1]))
		//spazi in mezzo
		fmt.Print(strings.Repeat(" ", spaziInMezzo))
		//carattere
		fmt.Print(string(stringa[len(stringa)-i-1]))
		//a capo
		fmt.Println()
		spaziInMezzo += 2

	}

}
