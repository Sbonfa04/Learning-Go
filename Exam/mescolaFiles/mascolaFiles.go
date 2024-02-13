/*
Scrivere un programma `mescolaFiles.go` che legge due file di input (i cui nomi sono passati su linea di
comando) e produce in output una "mescola" che prende una parola dal primo file e una dal secondo
(scrivendole seguite da un newline ciascuna) fino ad esaurire entrambi i file.
Se i file contengono numeri di parole diversi il programma deve comunque arrivare ad esaurire tutto
l'input (usando le parole rimanenti del file che ne contiene di più).

Se mancano parametri a linea di comando il programma termina col messaggio "Inserire DUE nomi di file".

Esempio, dati due file che contengono rispettivamente:
1) uno due tre quattro cinque
2) alpha beta gamma

L'output che viene generato è:
uno
alpha
due
beta
tre
gamma
quattro
cinque
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func stampa(file1 []string, file2 []string) {
	if len(file1) < len(file2) {
		file1, file2 = file2, file1
	}
	for i := 0; i < len(file1); i++ {
		fmt.Println(file1[i])
		if i < len(file2) {
			fmt.Println(file2[i])
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Inserire DUE nomi di file")
		return
	}
	var lista1, lista2 []string
	file1 := os.Args[1]
	file2 := os.Args[2]

	f1, _ := os.Open(file1)
	scanner := bufio.NewScanner(f1)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		lista1 = append(lista1, scanner.Text())
	}
	f1.Close()

	f2, _ := os.Open(file2)
	scanner2 := bufio.NewScanner(f2)
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		lista2 = append(lista2, scanner2.Text())
	}
	f2.Close()

	stampa(lista1, lista2)
}
