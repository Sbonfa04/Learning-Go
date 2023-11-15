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
	var altezza [200]int //nome, costante, tipo
	for i := 0; i < n; i++ {
		fmt.Scan(&altezza[i]) //per valori di n maggiori di 200 il programma va in panico
	}
	s := 0
	for i := 0; i < n; i++ {
		s += altezza[i] //calcolo la somma delle altezze
	}
	/*potrei utilizzare un for range ma vorrebbe  dire sommare
	inutilmente degli zeri fino a esaurimento array (in caso non
	ci fossero dati per ogni valore di array)*/
	media := float64(s) / float64(n)
	sq := 0.0
	for i := 0; i < n; i++ {
		sq += (float64(altezza[i]) - media) * (float64(altezza[i]) - media) //scarto quadratico
	}
	/*non si può sostituire con un for range perché i for range leggono
	tutto l'array mentre i for ternari leggono fino a n-1. (una volta
	finiti i valori il for range sommerebbe media^2 intuilmente fino a
	esaurimento array)*/
	sqm := math.Sqrt(sq / float64(n)) //scarto quadratico medio
	fmt.Println("media: ", media, "\n", "scarto quadratico medio: ", sqm, "\n")
}
