/**
N.B. leggere il file README.txt per istruzioni di compilazione, test e consegna

=== Cerchi ===

Scrivere un programma 'cerchi.go' dotato di:

- una struttura 'Cerchio' con campi:
	nome string
	x,y,raggio float64
		(dove x e y sono le coordinate del centro)
  dichiarati in quest'ordine

- una funzione 'newCircle(descr string) Cerchio'
	che data una stringa che descrive un cerchio
	(nome, coordinate x e y del centro, raggio , in quest'ordine e separati da spazi)
	restituisce il cerchio corrispondente

- una funzione 'distanzaPunti(x1,y1,x2,y2 float64) float64
	che restituisce la distanza tra i punti di coordinate (x1,y1) e (x2,y2)

- una funzione 'tocca(cerc1, cerc2 Cerchio) bool'
	che restituisce true se i due cerchi sono tangenti
	(internamente o esternamente); false altrimenti.
	Trattandosi di valori float, consideriamo una distanza
	trascurabile se è inferiore a 10^-6 (1e-6)

- una funzione maggiore(cerc1, cerc2 Cerchio) bool
	che restituisce true se cerc1 è più grande di cerc2;
	false altrimenti.

- una funzione main()
	che legge da standard input una sequenza (terminata da ctrl D)
	di righe che descrivono cerchi, del tipo:

uno 10.3 12.7 45.8
due 1.3 2.9 5.8
pippo 7.3 22.5 6.6

	ciascuna contenente nome, coordinate del centro e raggio di un
	cerchio, in quest'ordine.

Il programma crea per ciascuna riga il cerchio corrispondente, lo stampa.
Inoltre stampa se il cerchio, rispetto a quello precedente, è tangente o no e maggiore o no.
Il primo confronto è fatto rispetto al cerchio "zero" ("", 0, 0, 0).
Non sono richiesti controlli sulla correttezza dei dati in input, potete assumere che siano
sempre nell'ordine e del tipo previsto.




Esempio
-------
(vengono marcate con '>' le righe digitate da tastiera,
le altre sono l'output del programma)

>uno 0.5 0 2.5
{uno 0.5 0 2.5} non tangente, maggiore
>due 0 0 3
{due 0 0 3} tangente, maggiore
>tre 10.2 -8.4 1.5
{tre 10.2 -8.4 1.5} non tangente, minore o uguale

*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cerchio struct {
	nome         string
	x, y, raggio float64
}

func newCircle(descr string) Cerchio {
	descrizione := strings.Split(descr, " ")
	nome := descrizione[0]
	x, _ := strconv.ParseFloat(descrizione[1], 64)
	y, _ := strconv.ParseFloat(descrizione[2], 64)
	raggio, _ := strconv.ParseFloat(descrizione[3], 64)
	return Cerchio{nome, x, y, raggio}
}

func distanzaPunti(x1, y1, x2, y2 float64) float64 {
	distanza := math.Sqrt(((x1 - x2) * (x1 - x2)) + ((y1 - y2) * (y1 - y2)))
	return distanza
}

func tocca(cerc1, cerc2 Cerchio) bool {
	if distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)*1e-6 <= (cerc1.raggio+cerc2.raggio)*1e-6 {
		return true
	} else {
		return false
	}
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	if cerc1.raggio > cerc2.raggio {
		return true
	} else {
		return false
	}
}

func main() {
	var k = 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		riga := scanner.Text()

		var descrizioneZero string

		if k == 0 {
			descrizioneZero = "zero 0 0 0"
		}

		descrizione := strings.Join(strings.Split(string(riga), " "), " ")

		fmt.Print(newCircle(descrizione))

		if tocca(newCircle(descrizioneZero), newCircle(descrizione)) {
			fmt.Print(" tangente, ")
		} else {
			fmt.Print(" non tangente, ")
		}

		if maggiore(newCircle(descrizione), newCircle(descrizioneZero)) {
			fmt.Print("maggiore")
		} else {
			fmt.Print("minore o uguale")
		}

		fmt.Println()

		descrizioneZero = descrizione
		k++

	}
}
