/*es media e varianza con gli append*/

/*Dopo aver acquisito il numero di persone totali, ne calcola
la media dell'altezza*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var altezza []int

	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		altezza = append(altezza, x)
	}

	/*con la funzione append diamo la possibilità di aggiungere
	elementi finché è necessario senza indicare in anticipo
	la dimensione dell'array*/

	/*append inserisce i dati che ci mancano negli spazi liberi
	della slice finché non finisce. Qaundo la slice finisce viene
	creato nella heap un array identico a quello della slice ma
	con molti spazi liberi da poter utilizzare per futuri append*/

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
