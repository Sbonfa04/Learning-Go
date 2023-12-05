/*
Scrivere un programma lunghezze.go che legge riga per riga un testo da standard input (potete usare la ridirezione), terminato da EOF, e stampa quante parole ci sono di lunghezza 1, quante di lunghezza 2, ecc.
In particolare sperimentare quattro tipi di stampa:
una stampa della mappa
una stampa degli elementi (le coppie k:v) della mappa, non importa in che ordine (for range)
una stampa degli elementi in ordine dalla lunghezza minima a quella massima, comprese le lunghezze eventualmente non presenti nella mappa (che tipo di for?)
una stampa degli elementi in ordine dalla lunghezza minima a quella massima, escluse le lunghezze non presenti nella mappa (come evito di stampare le lunghezze che hanno 0 parole associate?).
*/
package main

import (
	"bufio"
	"fmt"

	//"fmt"
	//"strconv"
	//"math"
	"os"
)

func main() {

	m := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		parola := scanner.Text()
		m[len(parola)]++
	}

	fmt.Println(m)

	max := 0

	for k, v := range m {
		if max < k {
			max = k
		}
		fmt.Println(k, v)
	}

	for k := 0; k <= max; k++ {
		fmt.Println(k, m[k])
	}

	for k := 0; k <= max; k++ {
		if m[k] != 0 {
			fmt.Println(k, m[k])
		}
	}
}
