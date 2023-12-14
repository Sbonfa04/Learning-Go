/*
Scrivere un programma rettangolo.go, dotato di:
- una struttura Rettangolo, con campi base e altezza, in quest'ordine, di tipo int
- un metodo String() per Rettangolo che restituisce il disegno del rettangolo, pieno
e disegnato con '.', terminato da un new-line ('\n'). Ad esempio, per un Rettangolo di
base 5 e altezza 3, il metodo deve restituire la seguente stringa:
.....
.....
.....
Se però la base o l'altezza del rettangolo sono uguali a 0, il metodo String deve
restituire il messaggio "rettangolo degenere".
- una funzione main() che, dati due numeri naturali come
argomenti sulla linea di comando, emetta nel flusso d'uscita il disegno del rettangolo
che ha quei valori per la base e l'altezza (utilizzando il metodo String).
*/
package main

import (
	"fmt"
)

type Rettangolo struct {
	base, altezza int
}

func (r Rettangolo) String() string {

	var s string
	if r.base == 0 || r.altezza == 0 {
		return "rettangolo degenere"
	} else {
		for i := 0; i < r.altezza; i++ {
			for j := 0; j < r.base; j++ {
				s = s + "."
			}
			s = s + "\n"
		}
		return s
	}
}

func main() {

	var r Rettangolo
	fmt.Scan(&r.base, &r.altezza)
	fmt.Println(r.String())
}
