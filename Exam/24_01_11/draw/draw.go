package main

import (
	"fmt"
	"os"
	"strconv"
)

func DrawPoint(c byte, k int) string {
	var punto string
	for i := 0; i < k; i++ {
		punto = " "
	}
	return punto + string(c)
}

func DrawSegment(c byte, k, l int) string {
	var segmento, segmento2 string
	for i := 0; i < k; i++ {
		segmento = " "
	}
	for i := 0; i < l; i++ {
		segmento2 = string(c)
	}
	return segmento + segmento2
}

func main() {

	lun := os.Args[1]
	num := os.Args[2]
	char := os.Args[3]

	l, _ := strconv.Atoi(lun)
	n, _ := strconv.Atoi(num)
	c := char[1]

	if l < 0 && n < 0 {
		return
	} else {
		for i := 0; i < l; i++ {
			fmt.Print(DrawPoint(c, n))
		}
		fmt.Print(DrawSegment(c, n, l))
		for i := 0; i < l; i++ {
			for i := 0; i < l; i++ {
				fmt.Print(" ")
			}
			fmt.Print(DrawPoint(c, n))
		}
	}

}
