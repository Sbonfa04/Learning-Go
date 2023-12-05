package main

import "fmt"

func main() {
	const K = 10
	var prev, currm int
	var ascending bool

	for i := 0; i < K; i++ {

		prev = currm
		fmt.Scan(&currm)

		if currm < prev {
			ascending = false
			break
		} else {
			ascending = true
		}
	}

	if !ascending {
		fmt.Print("non ")
	}
	fmt.Println("crescente")
}
