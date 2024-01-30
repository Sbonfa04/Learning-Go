/*
Lessico
-------

Scrivere un programma lessico.go che legge da standard input un testo e registra, per ogni parola,
i numeri di riga in cui a parola stessa appare.

In particolare il programma lessico.go stampa (una volta sola) il seguente menu di opzioni:
+ (memorizza)
? (indica n. di riga)
p (stampa tutto)

poi legge un numero arbitrario di righe da standard input che iniziano tutte con uno di tre possibili
caratteri ('+', '?', 'p'), dove:

- le righe che iniziano con '+' sono il testo vero e proprio
- le righe che iniziano con '?' servono per fare una richiesta
- le righe 'p' servono per ottenere la stampa

e per ogni tipo di riga fa quanto indicato più sotto.

Il programma termina quando riceve un "end of file" (cioè EOF, pressione da tastiera di 'CTRL-d' dopo
aver dato <invio>)

Se la riga inizia con:
- "+ ", il programma considera il testo che segue il carattere "+ " e memorizza le parole
(stringhe separate da white space) che lo costituiscono, insieme ai numeri di riga in cui la
parola è apparsa nel testo (la prima riga  che inizia con '+' è la numero 1, e vanno considerate
solo le righe che iniziano con '+');
- "? ", il programma considera la parola o la frase che segue "? " e stampa
- "stringa:" seguito da uno spazio e dalla stringa (parola o frase) fornita
- "righe:" seguito da uno spazio e dalla lista dai numeri di riga in cui è comparsa quella stringa
  nel testo letto.

Se la riga è costituita da
- "p", il programma stampa le parole memorizzate finora, in ordine lessicografico, una per riga,
ciascuna seguita da ':', spazio, e dall’elenco dei numeri di riga in cui la parola stessa è comparsa
nel testo letto (vedi esempio).

Il programma deve essere dotato di una funzione

	func stampa(m map[string][]int)

che, data una mappa m, stampa in ordine lessicografico, una per riga, le chiavi della mappa,
ciascuna seguita, sulla stessa riga, da ": " e dal suo valore (vedi Esempio sotto)

Esempio di esecuzione
---------------------

$ ./lessico < medio.input
+ (memorizza)
? (indica n. di riga)
p (stampa tutto)
stringa: ha
righe [1]
befana: [1]
e: [1]
fazzoletto: [1]
gonna: [1]
ha: [1]
il: [1]
la: [1 1]
rattoppata: [1]
stringa: la
righe [1 1 3]
stringa: befana
righe [1 3]
stringa: ha il
righe []
befana: [1 3]
e: [1]
fazzoletto: [1]
gonna: [1]
ha: [1]
il: [1]
la: [1 1 3]
ma: [2]
poverina: [2]
quest'anno: [2]
raffreddata: [3]
rattoppata: [1]
è: [3]
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stampa(m map[string][]int) {
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func main() {
	fmt.Print("+ (memorizza)\n? (indica n. di riga)\np (stampa tutto)\n")
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string][]int)
	var nRiga int = 1
	for scanner.Scan() {
		riga := scanner.Text()
		rigaSep := strings.Split(riga, " ")

		if rigaSep[0] == "+" {
			// memorizza
			for _, parola := range rigaSep[1:] {
				m[parola] = append(m[parola], nRiga)
			}
			nRiga++
		} else if rigaSep[0] == "?" {
			// indica n. di riga
			parola := rigaSep[1]
			fmt.Printf("stringa: %s\n", parola)
			fmt.Printf("righe: %v\n", m[parola])
		} else if rigaSep[0] == "p" {
			// stampa tutto
			stampa(m)
		}
	}
}
