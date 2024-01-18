/*
Scrivere un programma carrello.go che simula gli spostamenti di un carrello lungo un percorso lineare a celle unitarie, in cui sono collocati degli oggetti di peso specificato.
Il carrello deve avanzare di una cella alla volta e ripulire il percorso caricando gli oggetti man mano che li incontra, ma senza superare il suo carico massimo.
Il carrello, quando si trova su una cella in cui c'è un oggetto, deve caricarlo; se non può caricarlo (il controllo lo fa sempre e solo mentre è sulla cella precedente), non deve avanzare su quella cella.
Quando l'oggetto sulla cella successiva farebbe superare il carico massimo, il carrello, invece di procedere, deve andare all'inizio del percorso, scaricare lì (nella cella 0) tutto il suo carico, e ripartire a pulire il percorso dagli oggetti che ancora vi si trovano.
Una volta caricato l'ultimo oggetto e ripulito tutto il percorso, il carrello deve svuotare il suo carico nella cella 0 e portarsi sulla cella 1.

Il programma deve essere dotato di:

- una struttura Carrello con campi
	- caricoMax int
	- posizione int
	- carico int
  dichiarati in quest'ordine

- un metodo String() string
  per Carrello che restituisce una descrizione del carrello che specifica posizione, carico e carico massimo, nel formato:
  posizione <posizione>, carico <carico> (max <caricoMax>)
  Ad esempio, per un Carrello con carico massimo 50 kg, che si trova sulla cella 25 e sta trasportando 42 kg, restituisce la stringa:
  "posizione 25, carico 42 (max 50)"

- una funzione aggiornaStato(c *Carrello, posizione, carico int) bool
  che aggiorna i dati del carrello (numero di cella e carico) e restuisce true se posizione e carico non sono negativi e se il carico non supera il carico massimo; altrimenti non fa nessun aggiornamento e restituisce false.

- una funzione main() che legge da file, il cui nome è passato su linea di comando, la descrizione di un percorso (a celle) con oggetti da rimuovere (vedi esempio sotto).
Una cella è delimitata da '|' e le celle sono numerate a partire da 0.
Se una cella se è vuota (se c'è uno spazio bianco tra '|'), in quella posizione non c'è nessun oggetto; se invece la cella contiene un numero, in quella posizione c'è un oggetto di quel peso.

Ad esempio, la seguente è la descrizione di un percorso:

"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"

che ha un oggetto di peso 12 nella cella 3, uno di peso 4 nella cella 4, ecc. come mostrato qui sotto:

"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"
  0 1 2 3  4 5 6 7  8 9 10 .....

La funzione main deve istanziare un Carrello che porta un massimo di 15 kg, si trova nella cella 0 del percorso, e ha un carico pari a 0, e poi gestire il Carrello in modo che ripulisca il percorso letto da file, operando come descritto sopra.
Si può dare per scontato, senza fare controlli, che sul percorso non ci siano pesi maggiori del peso massimo che il carrello può caricare.

Il main deve stampare:
- il percorso come è prima di iniziare il lavoro di pulizia;
- il carrello ogni volta che si verifica il caso che non è possibile caricare un nuovo oggetto;
- il percorso dopo ogni scarico;
- il carrello alla fine del lavoro;
- il numero di viaggi indietro fatti per rimuovere tutti gli oggetti dal percorso;
- il massimo peso trovato lungo il percorso;
- per ogni peso (in ordine crescente di peso), il numero di oggetti trovati di quel peso.

(Nota. si consiglia in fase di sviluppo di stampare anche la situazione del carrello dopo ogni carico)

Nel caso manchi il nome del file, il programma deve stampare il messaggio "manca nome file" e terminare.

Esempio di esecuzione per il percorso dato sopra
------------------------------------------------

| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 3, carico 12 (max 15)
|12| | | |4| | | |10| | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 12, carico 14 (max 15)
|26| | | | | | | | | | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 19, carico 9 (max 15)
|35| | | | | | | | | | | | | | | | | | | |12| | | | |3| |
|50| | | | | | | | | | | | | | | | | | | | | | | | | | |
carrello: posizione 1, carico 0 (max 15)
n viaggi: 4
peso max: 12
oggetti trovati:
1 ogg. di peso 3
2 ogg. di peso 4
1 ogg. di peso 5
1 ogg. di peso 10
2 ogg. di peso 12



*****************************************************************
>>> NOTA BENE

I test qui sotto funzionano correttamente su Linux e con
l'oracolo fornito, ma dato che l'oracolo è compilato per Linux
tutti i test basati su di esso non funzioneranno (falliranno)
su altre piattaforme.
Ergo sono stati forniti anche dei test basati sui
file ".expected" per chi usa Windows o Mac a casa.
*****************************************************************

*/

package main

import (
	"fmt"
	"testing"
)

var prog = "./carrello"
var diffwidth = 120

func TestEsisteCarrello(t *testing.T) {
	var c Carrello
	fmt.Println(c)
}

func TestString1(t *testing.T) {
	var c Carrello
	s := fmt.Sprint(c)
	if s != "posizione 0, carico 0 (max 0)" {
		t.Fail()
	}
}

func TestString2(t *testing.T) {
	c := Carrello{10, 3, 6}
	s := fmt.Sprint(c)
	if s != "posizione 3, carico 6 (max 10)" {
		t.Fail()
	}
}

func TestAggiorna(t *testing.T) {
	c := Carrello{10, 0, 0}
	var ok bool

	aggiornaStato(&c, 5, 8)
	if c.pos != 5 || c.carico != 8 {
		fmt.Println("non aggiorna correttamente (5,8)")
		t.Fail()
	}

	ok = aggiornaStato(&c, 8, 15)
	if ok {
		fmt.Println("non aggiorna correttamente (8,15)")
		t.Fail()
	}

	aggiornaStato(&c, 0, 0)
	if c.pos != 0 || c.carico != 0 {
		fmt.Println("non aggiorna correttamente (0,0)")
		t.Fail()
	}

	ok = aggiornaStato(&c, -2, 3)
	if ok {
		fmt.Println("non controlla i parametri")
		t.Fail()
	}

	ok = aggiornaStato(&c, 2, -3)
	if ok {
		fmt.Println("non controlla i parametri")
		t.Fail()
	}
}

func TestPercorso(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"percorsoEsempio.input")
}

/* */
func TestConExpectedPercorso(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"percorsoEsempio.expected",
		"percorsoEsempio.input")
}

/**/

func TestPercorsoBreve(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"percorsoBreve.input")
}

/**/
func TestConExpectedPercorsoBreve(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"percorsoBreve.expected",
		"percorsoBreve.input")
}

/**/

/*
func TestTODO(t *testing.T) {
	fmt.Println("********************** AGGIUNGERE TEST? ********************** ")
	//t.Fail()
}
*/
