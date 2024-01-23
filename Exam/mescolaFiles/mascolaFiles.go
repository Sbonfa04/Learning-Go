/*
Scrivere un programma `mescolaFiles.go` che legge due file di input (i cui nomi sono passati su linea di comando) e produce in output una "mescola" che prende una parola dal primo file e una dal secondo (scrivendole seguite da un newline ciascuna) fino ad esaurire entrambi i file.
Se i file contengono numeri di parole diversi il programma deve comunque arrivare ad esaurire tutto l'input (usando le parole rimanenti del file che ne contiene di più).

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
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Inserire DUE nomi di file")
		return
	}

	file1, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}
	defer file1.Close()

	file2, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}
	defer file2.Close()

	var riga1, riga2 string
	var eof1, eof2 bool

	for {
		if !eof1 {
			_, err = fmt.Fscanf(file1, "%s", &riga1)
			if err != nil {
				eof1 = true
			}
		}

		if !eof2 {
			_, err = fmt.Fscanf(file2, "%s", &riga2)
			if err != nil {
				eof2 = true
			}
		}

		if eof1 && eof2 {
			break
		}

		if !eof1 {
			fmt.Println(riga1)
		}

		if !eof2 {
			fmt.Println(riga2)
		}
	}
}
