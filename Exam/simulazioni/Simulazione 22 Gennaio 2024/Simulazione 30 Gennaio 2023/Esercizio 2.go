package main

import (
	"fmt"
	"os"
	"sort"
	"unicode"
)

func GeneraSottosequenze(sequenza []string) (sottosequenze map[string]int) {

	sottosequenze = make(map[string]int)
	tempSubString := ""

	for i := 0; i < len(sequenza); i++ {
		tempSubString = string(unicode.ToLower(rune(sequenza[i][0])))

		for j := i + 1; j < len(sequenza); j++ {
			tempSubString += string(unicode.ToLower(rune(sequenza[j][0])))

			if tempSubString[0] == tempSubString[len(tempSubString)-1] { //se la sottostringa inizia e finisce con la stessa lettera
				_, check := sottosequenze[tempSubString]

				if check {
					sottosequenze[tempSubString]++
				} else {
					sottosequenze[tempSubString] = 1
				}
			}
		}

	}

	return
}

func main() {
	sottoSequenze := GeneraSottosequenze(os.Args[1:])

	keys := make([]string, len(sottoSequenze))

	i := 0
	for key, _ := range sottoSequenze {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, sottoSequenze[key])
	}
}
