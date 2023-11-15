/*Dopo aver acquisito il numero di persone totali, ne calcola
la media dell'altezza*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)
	var altezza []int          //dichiarazione slice
	altezza = make([]int, 100) //creazione della slice
	for i := 0; i < n; i++ {
		fmt.Scan(&altezza[i]) //per valori di n maggiori di 200 il programma va in panico
	}
	s := 0
	for _, x := range altezza {
		s += x //calcolo la somma delle altezze
	}

	media := float64(s) / float64(n)
	sq := 0.0
	for _, x := range altezza {
		sq += (float64(x) - media) * (float64(x) - media) //scarto quadratico
	}

	sqm := math.Sqrt(sq / float64(n)) //scarto quadratico medio
	fmt.Println("media: ", media, "\n", "scarto quadratico medio: ", sqm, "\n")
}
