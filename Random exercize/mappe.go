/*
- Scrivere una funzione che dato il nome di un file ne restituisce una mappa
float64 che ha come chiavi le parole presenti nel file e come valori il numero
di volte che la parola compare nel file.
- Scrivere una funzione che date due mappe come sopra dice se sono più o
meno uguali
- Scrivere un programma che prende sulla riga di comando due nomi di due file e
tenta di capire se sono scritti nella stessa lingua
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordCount(filename string) (map[string]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	freq := make(map[string]float64)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		freq[word]++
	}

	return freq, nil
}

func compareMaps(map1, map2 map[string]float64) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || value1 != value2 {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two filenames")
		os.Exit(1)
	}

	map1, err := wordCount(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	map2, err := wordCount(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if compareMaps(map1, map2) {
		fmt.Println("The files are written in the same language")
	} else {
		fmt.Println("The files are written in different languages")
	}
}
