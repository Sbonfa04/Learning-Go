package main

import (
	"testing"
)

var prog = "./max"
var diffwidth = 120

/*
func TestTODO(t *testing.T) {
	fmt.Println("RIMPOLPARE TEST!!!", prog)
	t.Fail()
}
*/

func TestRico1(t *testing.T) {
	l := []int{3, 5, 6, 8, 9, 9, 10000, 5, 6, 6, 6, 5, 5}
	if recursiveMax(l) != 10000 {
		t.Fail()
	}
}

func TestRico2(t *testing.T) {
	l := []int{3324, 5, 6, 86788, 9, 9, 10000, 5, 6, 424236, -6666, -5, 5}
	if recursiveMax(l) != 424236 {
		t.Fail()
	}
}

func TestRicoZero(t *testing.T) {
	l := []int{0}
	if recursiveMax(l) != 0 {
		t.Fail()
	}
}

func TestRicoNeg(t *testing.T) {
	l := []int{-60, -89, -343434, -8}
	if recursiveMax(l) != -8 {
		t.Fail()
	}
}

func TestMainMinimale(t *testing.T) {
	/* - una funzione `main()` che legga da standard input (ctrl D per terminare) una lista di numeri interi (che posono essere positivi, negativi, nulli) ed emetta nel flusso d'uscita il massimo tra i numeri letti.
	 */
	LanciaGenerica(t,
		prog,
		"-60 -89 -343434 -8 12 567 76",
		"567\n")
}

func TestMain(t *testing.T) {
	/* - una funzione `main()` che legga da standard input (ctrl D per terminare) una lista di numeri interi (che posono essere positivi, negativi, nulli) ed emetta nel flusso d'uscita il massimo tra i numeri letti.
	 */
	LanciaGenerica(t,
		prog,
		"7165 -60 -89 -343434 -8 12 567 76 298371937182",
		"298371937182\n")
}
