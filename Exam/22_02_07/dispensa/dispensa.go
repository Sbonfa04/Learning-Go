package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var dispensa = make(map[string]float64)
	if len(os.Args) == 2 {
		if AggiornaDispensa(dispensa, os.Args[1]) {
			fmt.Println("file corretto")
		} else {
			fmt.Println("file non corretto")
		}
	} else {
		AggiornaDispensa(dispensa, os.Args[1])
		for i := 2; i < len(os.Args); i++ {
			fmt.Println(os.Args[i], Rimanenza(dispensa, os.Args[i]))
		}
	}
}

func AggiornaDispensa(dispensa map[string]float64, filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			riga := scanner.Text()
			rigasplit := strings.Split(riga, ",")
			if rigasplit[0] == "approv" {
				valore, _ := strconv.ParseFloat(rigasplit[2], 64)
				dispensa[rigasplit[1]] += valore
			} else if rigasplit[0] == "uso" {
				valore, _ := strconv.ParseFloat(rigasplit[2], 64)
				dispensa[rigasplit[1]] -= valore
				if dispensa[rigasplit[1]] < 0 {
					return false
				}
			}
		}
		return true
	}
}

func Rimanenza(dispensa map[string]float64, alimento string) float64 {
	for i, quantità := range dispensa {
		if i == alimento {
			return quantità
		}
	}
	return 0
}
