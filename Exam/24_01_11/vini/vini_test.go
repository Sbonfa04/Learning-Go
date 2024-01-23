/*
Vini
----

Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

  - una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
    nome string
    anno int
    gradi float32
    cl int

  - una funzione
    CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool)
    che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a questi dati e la restituisce insieme al valore 'true',
    altrimenti restituisce una bottiglia "zero-value" e 'false'.

  - una funzione
    CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
    che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo CSV (vedi esempio nelle specifiche del main);
    se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e la restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
    Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

  - un **metodo** per BottigliaVino
    String() string
    che restituisce una riga di descrizione della bottiglia nel seguente formato:  <nome>, <anno>, <gradi>°, <cl>cl
    (cioè ad es. "Rasol, 2018, 14°, 750cl" per la prima riga dell'esempio sopra).
    Suggerimento: i "format verb" %g e %v formattano i float omettendo il punto decimale quando non ci sono cifre dopo la virgola

- una funzione main() che legge da un file (il cui nome è passato da linea di comando) delle righe che contengono ognuna i dati relativi ad una bottiglia di una cantina, separati da virgole, nel formato del seguente esempio (nome,anno,gradi,cl):

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

e stampa su stdout:

  - l'elenco delle bottiglie della cantina (esattamente nello stesso formato rappresentato qui sopra).
    Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

  - il numero di bottiglie nella cantina

  - la bottiglia con gradazione massima

  - la bottiglia più vecchia

  - i cl totali di vino della cantina

Esempio di esecuzione con input vini.input
---------------------

Rasol, 2018, 14°, 750cl
Camunnorum, 2015, 15°, 750cl
Dom Perignon, 2019, 12.5°, 1500cl
Balon, 2013, 15°, 750cl
Verdicchio, 2020, 11°, 375cl
Rasol, 2018, 14°, 1000cl
Verdicchio, 2020, 11°, 375cl
n. bottiglie: 7
bottiglia con grado max: Camunnorum, 2015, 15°, 750cl
bottiglia più vecchia: Balon, 2013, 15°, 750cl
tot vino: 5500 cl
*/
package main

import (
	//"strings"
	//"log"
	//"fmt"
	//"os/exec"
	//"strings"

	"testing"
)

func TestCreaBottiglia(t *testing.T) {
	b, correct := CreaBottiglia("Dom Perignon", 2019, 12.5, 1500)
	if !correct {
		t.FailNow()
	}
	if b.nome != "Dom Perignon" {
		t.FailNow()
	}
	if !intorno(b.gradi, 12.5) {
		t.FailNow()
	}
	//fmt.Println(b)
}

func TestCreaBottiglieSbagliate(t *testing.T) {
	_, correct := CreaBottiglia("", 2019, 12.5, 1500)
	if correct {
		t.FailNow()
	}
	_, correct = CreaBottiglia("Dom", 2000, -12.5, 1500)
	if correct {
		t.FailNow()
	}
	_, correct = CreaBottiglia("Dom", 2000, 12.5, -1500)
	if correct {
		t.FailNow()
	}

	_, correct = CreaBottiglia("Dom", 0, 12.5, 1500)
	if correct {
		t.FailNow()
	}

	_, correct = CreaBottigliaDaRiga("Verdicchio,2020,11")
	if correct {
		t.FailNow()
	}
}

func TestCreaBottigliaRiga(t *testing.T) {
	b, correct := CreaBottigliaDaRiga("Verdicchio,2020,11,375")
	if !correct {
		t.FailNow()
	}
	if b.nome != "Verdicchio" {
		t.FailNow()
	}
	if !intorno(b.gradi, 11.0) {
		t.FailNow()
	}
	if b.anno != 2020 {
		t.FailNow()
	}
	if b.cl != 375 {
		t.FailNow()
	}

	//fmt.Println(b)
}

func TestString(t *testing.T) {
	b, _ := CreaBottiglia("Dom Perignon", 2019, 12.5, 1500)
	//fmt.Println(b)
	if b.String() != "Dom Perignon, 2019, 12.5°, 1500cl" {
		t.FailNow()
	}
}

var prog = "./vini"
var diffwidth = 120

func TestMain(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"vini.input") // arg1
	/* OLD
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"Rasol, 2018, 14°, 750cl\nCamunnorum, 2015, 15°, 750cl\nDom Perignon, 2019, 12.5°, 1500cl\nBalon, 2013, 15°, 750cl\nVerdicchio, 2020, 11°, 375cl\nRasol, 2018, 14°, 1000cl\nVerdicchio, 2020, 11°, 375cl\nn. bottiglie: 7\nbottiglia con grado max: Camunnorum, 2015, 15°, 750cl\nbottiglia più vecchia: Balon, 2013, 15°, 750cl\ntot vino: 5500 cl")
	*/
}
