package main

import (
	"fmt"
)

func main() {

	var n, s, h int
	fmt.Scan(&h)

	for h > 0 {
		n++
		s += h //accumulatore
		fmt.Scan(&h)
	}
	media := float64(s) / float64(n)
	fmt.Println(media)
}

//in alternativa s può usare un break
//var n, s, h int
//	fmt.Scan(&h)
//
//	for h > 0 {
//		fmt.Scan(&h)
//		if h<=0 {
//			break
//		}
//		n++
//		s+=h
//	}
//	media := float64(s) / float64(n)
//	fmt.Println(media)
//}
