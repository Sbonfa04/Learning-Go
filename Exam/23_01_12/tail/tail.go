/*
Realizzare un programma Go (nome file 'tail.go') che implementi un semplice 'tail', comando Unix
che estrae le ultime N righe di un file di testo.

Il programma deve prendere due parametri:
- numero N di linee da estrarre
- nome del file da elaborare

e stampare su standard output le ultime N righe del file. Se il file ha meno di N righe, il programma stamperà tutte quelle che ci sono.

Nota bene: NON va implementato invocando da Go il comando 'tail' di sistema.

## Esempio esecuzione

(presuppone il 'tail.go' già compilato, 'uno.input' è presente in questa directory)

$ ./tail 3 uno.input
remaining essentially unchanged. It was popularised in the 1960s with the release
of Letraset sheets containing Lorem Ipsum passages, and more
recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

	righe, _ := strconv.Atoi(os.Args[1])
	f := os.Args[2]
	file, _ := os.Open(f)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := len(lines) - righe; i < len(lines); i++ {
		if i >= 0 {
			println(lines[i])
		}
	}
	file.Close()
}
