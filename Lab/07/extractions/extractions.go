/*Scrivere un programma extractions.go con due funzioni:
- estraiPari(in []int) (out []int) che prende una slice di interi e ne restituisce
una che contiene solo i numeri di quella in ingresso che sono numeri pari.

- rimuoviMultipli(m int, in []int) (out []int) che prende un intero e una slice di interi
e ne restituisce una senza i multipli del numero passato di quella in ingresso.
Es.: rimuoviMultipli(5, in) restituisce, a partire da in, una slice senza i multipli di 5.
*/

package main

func estraiPari(in []int) (out []int) {

	for i := 0; i < len(in); i++ {
		if in[i]%2 == 0 {
			out = append(out, in[i])
		}
	}

	return out
}

func rimuoviMultipli(m int, in []int) (out []int) {

	for i := 0; i < len(in); i++ {
		if in[i]%m != 0 {
			out = append(out, in[i])
		}
	}

	return out
}

func main() {

}
