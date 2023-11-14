package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 1000000; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Print(rand.Intn(10), " ")
	}
	fmt.Print("\n")
}
