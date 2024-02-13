package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Prodotto struct {
	nome   string
	codice string
}

type Magazzino map[Prodotto]int

type Prodotti []Prodotto

//funzioni per soddisfare l'interfaccia necessaria all'ordinamento.
func (p Prodotti) Len() int {
	return len(p)
}

func (p Prodotti) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Prodotti) Less(i, j int) bool {
	return p[i].codice < p[j].codice
}

func NuovoMagazzino() Magazzino {
	return make(Magazzino)
}

func StringProdotto(p Prodotto) string {
	return "Codice: " + p.codice + ", Prodotto: " + p.nome
}

func CreaProdotto(nome string, codice string) Prodotto {
	return Prodotto{nome: nome, codice: codice}
}

func NumeroProdotti(m Magazzino) int {
	return len(m)
}

func Quantità(m Magazzino, p Prodotto) int {
	return m[p]
}

func ModificaProdotto(m Magazzino, p Prodotto, variazione int) (Magazzino, bool) {
	value, present := m[p]
	boolean := true

	if present {
		if (value + variazione) > 0 {
			m[p] += variazione
		} else if (value + variazione) == 0 {
			delete(m, p)
		} else {
			boolean = false
		}
	} else {
		if variazione > 0 {
			m[p] = variazione
		}
	}

	return m, boolean
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	magazzino := NuovoMagazzino()
	boolean := true
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		singleFields := strings.Split(scanner.Text(), ";")
		prodotto := CreaProdotto(singleFields[1], singleFields[0])
		variazione, _ := strconv.Atoi(singleFields[2])

		magazzino, boolean = ModificaProdotto(magazzino, prodotto, variazione)
		if boolean == false {
			fmt.Println("Errore alla riga ", lineNumber, ",", StringProdotto(prodotto), ", Quantità: ", magazzino[prodotto], ", Richiesta: ", variazione)
			return
		}
	}

	fmt.Println("Il Magazzino contiene " + string(NumeroProdotti(magazzino)) + " prodotti.")

	//ordinamento
	prodotti := make(Prodotti, len(magazzino))

	i := 0
	for prodotto, _ := range magazzino {
		prodotti[i] = prodotto
		i++
	}

	sort.Sort(prodotti)

	for _, prodotto := range prodotti {
		fmt.Println(StringProdotto(prodotto), ", Quantità: ", magazzino[prodotto])
	}
}
