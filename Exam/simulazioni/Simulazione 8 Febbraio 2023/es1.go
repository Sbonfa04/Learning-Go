/*
ESERCIZIO 2
Solitamente, un numero romano presenta i simboli posizionati in ordine di valore, dal più grande al
più piccolo. Il valore di un numero romano si calcola solitamente per somma del valore dei simboli
che lo compongono. Per esempio, 2 equivale a II in numero romano (I+I).12 è scritto come XII, che
è semplicemente X+II. In modo simile, 27 è scritto come XXVII, chi è XX + V + II. Tuttavia unostesso simbolo non può essere ripetuto più di 3 volte, infatti IIII non è un numero romano ben
formattato. Per esprimere il valore 4, la giusta sintassi e IV pertanto, quando un simbolo con
valore minore si presenta prima di un simbolo con valore maggiore, il valore più piccolo deve
essere sottratto al valore più grande.
Esempi:
- IV = V - I = 5 - 1 = 4
- MCMXLIV = M + (M - C) + (L - X) + (V - X) = 1000 + (1000 - 100) + (50 - 10) +5 - 1) = 1944
Il simbolo I può comparire alla sinistra dei simboli V e X, ma una sola volta.ad esempio X equivale
a 9. IX non è un numero romano ben formattato, in quanto il numero romano corrispondente a 8 e
VIII. Il simbolo X può comparire alla sinistra dei simboli L e C, ma una sola volta. Ad esempio il
numero 99 corrisponde a XCX: la somma di XC (90) e IX (nove). Il numero romano iC non è ben
formattato infine il simbolo ci può comparire alla sinistra dei simboli D e M ma una sola volta. Ad
esempio, il numero 900 corrisponde a CM. I simboli V, L e D non possono mai essere posizionati
alla sinistra di simboli con valore più grande (VC ad esempio non è ben formattato in quanto il
numero 95 corrisponde a XCV: XC + V).
Si scrive un programma che legga da riga di comando un numero romano (tipo string) e che
stampi il suo valore in formato decimale. A tal fine si implementi la funzione Romano2decimale
(n string) int che prende in input un numero romano ben formattato (tipo string) e
restituisce il valore corrispondente in formato decimale (tipo int). Si assume che il numero romano
inserito da riga di comando sia ben formattato. Si può inoltre assumere che la stringa in input sarà
composta da soli caratteri ASCII.
Esempio d’esecuzione:
$ go run esercizio_2.go CXLI
141
$ go run esercizio_2.go MCMLXXXIII
1983
$ go run esercizio_2.go CCCLXXXIX
389
$ go run esercizio_2.go MMMCMXCIX
3999
$ go run esercizio_2.go MDCCLXXXIV
1784
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Romano2decimale(os.Args[1]))
}

func Romano2decimale(n string) int {
	romani := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	decimale := 0
	for i := 0; i < len(n); i++ {
		if i < len(n)-1 && romani[n[i]] < romani[n[i+1]] {
			decimale -= romani[n[i]]
		} else {
			decimale += romani[n[i]]
		}
	}
	return decimale
}
