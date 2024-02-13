package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func Sottostringhe(s string) map[string]int {
	mappaSottostringhe := make(map[string]int)

	for letteraIniziale := 0; letteraIniziale < len(s); letteraIniziale++ {
		for letteraControllo := letteraIniziale + 1; letteraControllo < len(s) && s[letteraControllo-1] < s[letteraControllo]; letteraControllo++ {
			_, check := mappaSottostringhe[s[letteraIniziale:letteraControllo+1]]
			if !check {
				mappaSottostringhe[s[letteraIniziale:letteraControllo+1]] = 1
			} else {
				mappaSottostringhe[s[letteraIniziale:letteraControllo+1]]++
			}
		}
	}
	return mappaSottostringhe
}

func onlyLowerCase(stringa string) bool {
	for _, v := range stringa {
		if !unicode.IsLower(v) || v > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		stringa := scanner.Text()
		if onlyLowerCase(stringa) {

			mappaSubstring := Sottostringhe(stringa)
			
			//ordinamento
			sliceKeys := make([]string, len(mappaSubstring))
			i := 0
			for chiave, _ := range mappaSubstring {
				sliceKeys[i] = chiave
				i++
			}

			sort.Strings(sliceKeys)

			for _, v := range sliceKeys {
				fmt.Println(v, mappaSubstring[v])
			}
		} else {
			return
		}
	}
}
