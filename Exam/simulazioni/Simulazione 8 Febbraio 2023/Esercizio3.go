package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prodotto struct {
	nome   string
	codice string
}

type Magazzino map[Prodotto]int

func NuovoMagazzino() Magazzino {
	magazzino := make(Magazzino)
	return magazzino
}

func StringProdotto(p Prodotto) string {
	return "Codice: " + p.codice + " Prodotto: " + p.nome
}

func CreaProdotto(nome string, codice string) Prodotto {
	return Prodotto{nome, codice}
}

func NumeroProdotti(m Magazzino) int {
	return len(m)
}

func Quantità(m Magazzino, p Prodotto) int {
	return m[p]
}

func ModificaProdotto(m Magazzino, p Prodotto, variazione int) (Magazzino, bool) {
	returnValue := true
	_, ok := m[p]
	if !ok { //se il prodotto non è in magazzino
		m[p] = variazione
	} else if m[p]+variazione > 0 {
		m[p] += variazione
	} else if m[p]+variazione == 0 {
		delete(m, p)
	} else {
		returnValue = false
	}
	return m, returnValue
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNumber := 0

	magazzino := NuovoMagazzino()

	for scanner.Scan() {
		lineNumber++
		codiceProdottoQuantità := strings.Split(scanner.Text(), ";")
		prodottoInserito := CreaProdotto(codiceProdottoQuantità[1], codiceProdottoQuantità[0])
		quantità, _ := strconv.Atoi(codiceProdottoQuantità[2])
		magazzino, ret := ModificaProdotto(magazzino, prodottoInserito, quantità)
		if !ret {
			fmt.Println("Errore alla riga ", lineNumber, StringProdotto(prodottoInserito), Quantità(magazzino, prodottoInserito), "Richiesta: ", codiceProdottoQuantità[2])
		}
	}

	for key, value := range magazzino {
		fmt.Println(StringProdotto(key), ",Quantità: ", value)
	}

}
