package main

import (
	"bufio"
	"os"
	//"strconv"
	//"math"
	//"bufio"
	//"os"
)

func contaVocali(s string, vocali map[rune]int) {

}

func main() {
	m := make(map[rune]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		vocale := scanner.Text()

	}
}
