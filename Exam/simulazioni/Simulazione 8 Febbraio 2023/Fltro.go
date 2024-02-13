package main

import (
	"fmt"
	"os"
)

func main() {
	stringa := os.Args[1]
	prev := ""

	for i, v := range stringa {
		fmt.Println(stringa[i:] + prev)
		prev += string(v)
	}
}
