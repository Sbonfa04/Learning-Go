package main

import (
	"fmt"
	"os"
	"strconv"
)

func GeneraNumeri(N, K int) (sottonumeri []int) {
	NArr := strconv.Itoa(N)

	for i := 0; i+K <= len(NArr); i++ {
		value, _ := strconv.Atoi(NArr[0:i] + NArr[i+K:])
		sottonumeri = append(sottonumeri, value)
	}

	return
}

func FiltraNumeri(sl []int) (result []int) {

	for _, v := range sl {
		sommaDivisori := 0
		//calcolo divisori
		for i := 1; i <= v/2; i++ {
			if v%i == 0 {
				sommaDivisori += i
			}
		}

		if v < sommaDivisori {
			result = append(result, v)
		}
	}

	return
}

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	K, _ := strconv.Atoi(os.Args[2])

	fmt.Println(GeneraNumeri(N, K))
	fmt.Println(FiltraNumeri(GeneraNumeri(N, K)))

}
