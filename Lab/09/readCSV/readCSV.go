/*

Scrivere un programma `readCSV.go` che legga da stdin una sequenza di dati in formato CSV
(Comma Separated Values) così organizzati:

	timestamp,temperatura,umidità

Dove:
- `timestamp` è una stringa nel formato YYYY-MM-DD/HH:mm (anno mese giorno ora minuto)
- `temperatura` è un float
- `umidità` è un float
- il separatore è una virgola ","

Il programma deve usare Scanf per leggere ogni riga dell'input.

Il programma deve calcolare il massimo e il minimo dei valori di temperatura e umidità e
stamparli con il timestamp (in formato diverso dall'originale, deve essere nella forma
"X(°/%) misura delle ore HH, minuti mm del giorno DD del mese MM dell'anno YYYY") in
cui sono stati misurati.

Ad esempio, il file fornito come input genera questo risultato:

minTemp: 1.0° misura delle ore 4, minuti 22 del giorno 11 del mese 12 dell'anno 2022
maxTemp: 15.0° misura delle ore 9, minuti 22 del giorno 11 del mese 12 dell'anno 2022
minHumid: 49.0% misura delle ore 9, minuti 24 del giorno 11 del mese 12 dell'anno 2022
maxHumid: 91.0% misura delle ore 9, minuti 31 del giorno 6 del mese 12 dell'anno 2022

*/

package main

import (
	"fmt"
)

func main() {

	var anno, mese, giorno, ora, minuto int
	var temperatura float64
	var umidità float64
	var maxHumid float64
	var maxTemp float64
	minHumid := 3242343242433242423432.32424
	minTemp := 324324323242424234324.3242
	var orarioMaxHumid []int
	var orarioMinHumid []int
	var orarioMaxTemp []int
	var orarioMinTemp []int

	for {
		_, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f\n", &anno, &mese, &giorno, &ora, &minuto, &temperatura, &umidità)
		if err != nil {
			break
		}

		if umidità < minHumid {
			minHumid = umidità
			orarioMinHumid = nil
			orarioMinHumid = append(orarioMinHumid, anno)
			orarioMinHumid = append(orarioMinHumid, mese)
			orarioMinHumid = append(orarioMinHumid, giorno)
			orarioMinHumid = append(orarioMinHumid, ora)
			orarioMinHumid = append(orarioMinHumid, minuto)
		}

		if umidità > maxHumid {
			maxHumid = umidità
			orarioMaxHumid = nil
			orarioMaxHumid = append(orarioMaxHumid, anno)
			orarioMaxHumid = append(orarioMaxHumid, mese)
			orarioMaxHumid = append(orarioMaxHumid, giorno)
			orarioMaxHumid = append(orarioMaxHumid, ora)
			orarioMaxHumid = append(orarioMaxHumid, minuto)
		}

		if temperatura < minTemp {
			orarioMinTemp = nil
			minTemp = temperatura
			orarioMinTemp = append(orarioMinTemp, anno)
			orarioMinTemp = append(orarioMinTemp, mese)
			orarioMinTemp = append(orarioMinTemp, giorno)
			orarioMinTemp = append(orarioMinTemp, ora)
			orarioMinTemp = append(orarioMinTemp, minuto)
		}

		if temperatura > maxTemp {
			orarioMaxTemp = nil
			maxTemp = temperatura
			orarioMaxTemp = append(orarioMaxTemp, anno)
			orarioMaxTemp = append(orarioMaxTemp, mese)
			orarioMaxTemp = append(orarioMaxTemp, giorno)
			orarioMaxTemp = append(orarioMaxTemp, ora)
			orarioMaxTemp = append(orarioMaxTemp, minuto)
		}

	}

	fmt.Printf("minTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", minTemp, orarioMinTemp[3], orarioMinTemp[4], orarioMinTemp[2], orarioMinTemp[1], orarioMinTemp[0])
	fmt.Printf("maxTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", maxTemp, orarioMaxTemp[3], orarioMaxTemp[4], orarioMaxTemp[2], orarioMaxTemp[1], orarioMaxTemp[0])
	fmt.Printf("minHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", minHumid, orarioMinHumid[3], orarioMinHumid[4], orarioMinHumid[2], orarioMinHumid[1], orarioMinHumid[0])
	fmt.Printf("maxHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", maxHumid, orarioMaxHumid[3], orarioMaxHumid[4], orarioMaxHumid[2], orarioMaxHumid[1], orarioMaxHumid[0])
}
