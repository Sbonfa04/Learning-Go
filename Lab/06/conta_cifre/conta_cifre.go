package main

import (
	"bufio"
	//"fmt"
	"os"
	//"strconv"
	//"math"
)

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	var c int
	//var cifre [10]int
	for _, i := range s {
		if i <= '0' || i >= '9' {
			c++
		}
	}
	if c > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		parola := scanner.Text()
		if parola == "stop" {
			break
		}
	}
}
