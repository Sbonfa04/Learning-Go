package main

import (
	"bufio"
	"fmt"
	"os"
)

type BottigliaVino struct {
	nome  string
	anno  int
	gradi float32
	cl    int
}

func CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool) {
	if nome != "" && anno > 0 && gradi > 0 && cl > 0 {
		return BottigliaVino{nome, anno, gradi, cl}, true 
	} else {
		return BottigliaVino{}, false
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	var b BottigliaVino
	var nome string
	var anno, cl int
	var gradi float32

	_, err

}

func (b BottigliaVino) String() string {
	
}

func main() {
	
}
