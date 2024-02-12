/*
Cancellazioni
=============

Scrivere un programma 'cancellazioni.go' dotato di:

- una funzione

 	func cancella(lista []string) []string

  che, per ogni numero n >0 (intero) presente in lista,
  cancella il numero stesso e gli n elementi successivi della lista
  (o quelli che ci sono per arrivare alla fine della lista)
  e restituisce la nuova lista così prodotta;

- una funzione main() che legge da file (il cui nome viene passato
  come parametro su linea di comando) un testo di parole (sequenze di caratteri separate da whitespace),
  tra cui anche numeri interi non negativi, stampa la lista delle parole lette e poi
  la nuova lista ottenuta cancellando, per ogni numero n presente nella lista originale,
  il numero stesso e gli n elementi successivi (vedi sopra).

Se il programma viene lanciato con un numero di argomenti diverso da 1,
deve terminare stampando "Fornire 1 nome di file".
Se invece il file non esiste, il programma deve terminare stampando "File non accessibile".

Il file può contenere un numero arbitrario di parole e numeri, disposti su un numero arbitrario di righe di testo, senza vincoli sul numero e tipo di caratteri whitespace che separano parole e numeri e sul numero di cifre di cui sono composti i numeri.

Si può assumere che il file non contenga numeri negativi, non occorre fare questo controllo.

Esempi di esecuzione
--------------------

$ ./cancellazioni uno.input
[uno due 2 tre quattro cinque 1 sei sette 5 otto nove dieci]
[uno due cinque sette]

$ ./cancellazioni due.input
[uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]
[uno due cinque sette nove dieci]

$ ./cancellazioni quattro.input
[uno 10 due tre quattro cinque sei sette otto nove dieci undici dodici]
[uno dodici]

$ ./cancellazioni
Fornire 1 nome di file

$ ./cancellazioni  FileNonEsistente.txt
File non accessibile

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cancella(lista []string) []string {
	var newLista []string
	for i := 0; i < len(lista); i++ {
		num, err := strconv.Atoi(lista[i])
		if err != nil {
			newLista = append(newLista, lista[i])
		} else {
			i += num
		}
	}
	return newLista
}

func main() {
	var lista []string
	if len(os.Args) < 2 {
		fmt.Println("Fornire 1 nome di file")
	} else {
		f := os.Args[1]
		file, err := os.Open(f)
		if err != nil {
			fmt.Println("File non accessibile")
			return
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			lista = append(lista, scanner.Text())
		}
		fmt.Println(lista)
		fmt.Println(cancella(lista))
	}
}
