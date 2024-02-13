package main

import (
	"math"
	"os"
	"strings"
)

func main() {
	parameter := []rune(os.Args[1])
	parameterLength := len(parameter)

	for index := 0; index < (parameterLength*2)-1; index += 2 {
		//print spaces before string
		k := int(math.Abs(float64((parameterLength - index - 1) / 2)))

		print(strings.Repeat(" ", k))

		//print runes
		print(string(parameter[:parameterLength-k*2]))

		//newline
		println("")
	}
}
