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
	"fmt"
	"testing"
	//"strings"
	//"log"
)

var prog = "./cancellazioni"

// var progname = "./cancellazioniAlternativo"
var diffwidth = 120

func TestMainUno(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"uno.input") // arg1
}

func TestMainDue(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"due.input") // arg1
}

func TestMainTre(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"tre.input") // arg1
}

func TestMainLungo(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"lungo.input") // arg1
}

/*
	-LanciaGenerica(t, progname, "NIENTE", "[uno due 2 tre quattro cinque 1 sei sette 5 otto nove dieci]\n[uno due cinque sette]\n", "uno.input")
	-LanciaGenerica(t, progname, "NIENTE", "[uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]\n[uno due cinque sette nove dieci]\n", "due.input")
	-LanciaGenerica(t, progname, "NIENTE", "[A b C d E f 2 A b C d E f A 3 b C d E f uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]\n[A b C d E f C d E f A E f uno due cinque sette nove dieci]\n", "tre.input")
	-LanciaGenerica(t, progname, "NIENTE", "[uno due 2 tre quattro cinque 1 sei sette 5 otto nove dieci uno due 11 tre quattro cinque sei sette otto nove dieci undici dodici tredici quattordici quindici sedici diciassette]\n[uno due cinque sette quattordici quindici sedici diciassette]\n", "lungo.input")
*/

func TestSituazioniFileNonEsistente(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"nulla") // arg1
	//LanciaGenerica(t, progname, "NIENTE", "File non accessibile\n", "nulla")
}

func TestSituazioniNoArgs(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"")
	//LanciaGenerica(t, progname, "NIENTE", "Fornire 1 nome di file\n")
}

func TestFunzione(t *testing.T) {
	fmt.Println("### Questo test verifica output atteso! (presuppone che il sorgente da testare sia già stato compilato)")

	var lista = []string{"uno", "due", "2", "tre", "quattro", "cinque"}
	fmt.Println(lista)

	//n:=fmt.Sprintf("%#v", cancella(lista))
	n := fmt.Sprint(cancella(lista))

	fmt.Println(n)

	if n != "[uno due cinque]" {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}
