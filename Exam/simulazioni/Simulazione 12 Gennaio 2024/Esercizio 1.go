package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcolaSoluzioni(a, b, c int) (soluzioni []float64) {
	delta := (b * b) - (4 * a * c)
	soluzione := (float64(-b) + math.Sqrt(float64(delta))) / float64(2*a)

	if delta == 0 {
		soluzioni = append(soluzioni, soluzione)
	} else if delta > 0 {
		soluzioni = append(soluzioni, soluzione)
		//seconda soluzione
		soluzione = (float64(-b) - math.Sqrt(float64(delta))) / float64(2*a)
		soluzioni = append(soluzioni, soluzione)
	}

	return
}

func main() {
	equazione := os.Args[1]                                //string
	sogliaMin, _ := strconv.ParseFloat(os.Args[2], 64)     //float
	sogliaEpsilon, _ := strconv.ParseFloat(os.Args[3], 64) //float

	singleFields := strings.FieldsFunc(equazione, func(x rune) bool { return x == '=' || x == '^' || x == 'x' })
	//indice 0 -> a
	//indice 1 -> b (primo carattere escluso)
	//indice 2 -> c
	a, _ := strconv.Atoi(singleFields[0])
	b, _ := strconv.Atoi(singleFields[1][1:])
	c, _ := strconv.Atoi(singleFields[2])

	soluzioni := calcolaSoluzioni(a, b, c)
	if len(soluzioni) == 0 {
		fmt.Println("Non ci sono soluzioni reali")
	} else if len(soluzioni) == 1 {
		fmt.Printf("Esiste un'unica soluzione reale: %.3f", soluzioni[0])
		if soluzioni[0]-math.Abs(float64(sogliaMin)) > float64(sogliaEpsilon) {
			fmt.Printf("\nLa soluzione %0.3f è maggiore della soglia.", soluzioni[0])
		}
	} else {
		fmt.Printf("Esistono due soluzioni reali: %.3f e %.3f", soluzioni[0], soluzioni[1])
		if soluzioni[0]-math.Abs(float64(sogliaMin)) > float64(sogliaEpsilon) {
			fmt.Printf("\nLa soluzione %0.3f è maggiore della soglia.", soluzioni[0])
		}
		if soluzioni[1]-math.Abs(float64(sogliaMin)) > float64(sogliaEpsilon) {
			fmt.Printf("\nLa soluzione %0.3f è maggiore della soglia.", soluzioni[1])
		}
	}
}
