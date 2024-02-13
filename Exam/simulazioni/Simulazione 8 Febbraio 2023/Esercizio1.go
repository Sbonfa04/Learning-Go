package main

import (
	"fmt"
)

func OnlyAscii(s string) bool {
	b := []byte(s) //len(b) vale 8 (è accentata è codificata su due byte)
	c := []rune(s) //nel caso di ciaoèè, len(c) vale 6

	if len(b) == len(c) {
		return true
	}
	return false
}

func main() {
	inputString := ""
	maxSub := make([]string, 0)
	tmpMaxSub := ""

	fmt.Println("Inserire input:")
	fmt.Scan(&inputString)

	if OnlyAscii(inputString) {
		fmt.Println(inputString)
	} else {
		for i, v := range inputString {
			if int(v) > 128 || i == len(inputString)-1 {
				if int(v) < 128 && i == len(inputString)-1 {
					tmpMaxSub += string(v)
				}

				if len(tmpMaxSub) > len(maxSub[len(maxSub)-1]) {
					maxSub = make([]string, 0)

					maxSub = append(maxSub, tmpMaxSub)
				} else if len(tmpMaxSub) == len(maxSub[len(maxSub)-1]) {
					maxSub = append(maxSub, tmpMaxSub)
				}
				tmpMaxSub = ""
			} else {
				tmpMaxSub += string(v)
			}
		}

		fmt.Println(maxSub)

	}

}
