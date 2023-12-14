package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("Gino.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var b []byte
	b = []byte("Sono Gino")
	n, err := f.Write(b)
	if err != nil {
		fmt.Println(err, "è ESPLOSO tutto")
		return
	}
	fmt.Println(n, "byte scritti")

}
