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
	var bottiglia BottigliaVino
	if nome == "" || anno <= 0 || gradi <= 0 || cl <= 0 {
		return bottiglia, false
	} else {
		bottiglia.nome = nome
		bottiglia.anno = anno
		bottiglia.gradi = gradi
		bottiglia.cl = cl
		return bottiglia, true
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	var bottiglia BottigliaVino
	var nome string
	var anno int
	var gradi float32
	var cl int
	_, err := fmt.Sscanf(riga, "%s,%d,%f,%d", &nome, &anno, &gradi, &cl)
	if err != nil {
		return bottiglia, false
	} else {
		return CreaBottiglia(nome, anno, gradi, cl)
	}
}

func (b BottigliaVino) String() string {
	return fmt.Sprintf("%s, %d, %v°, %dcl", b.nome, b.anno, b.gradi, b.cl)
}

func main() {
	var numBottiglie, maxAnno, clTot int
	var maxGradi float32
	var maxGradiBottiglia, maxAnnoBottiglia BottigliaVino
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riga := scanner.Text()
		if len(riga) > 0 {
			bottiglia, correct := CreaBottigliaDaRiga(riga)
			if correct {
				fmt.Println(bottiglia)
				numBottiglie++
				if bottiglia.gradi > maxGradi {
					maxGradi = bottiglia.gradi
					maxGradiBottiglia = bottiglia
				}
				if bottiglia.anno < maxAnno {
					maxAnno = bottiglia.anno
					maxAnnoBottiglia = bottiglia
				}
				clTot += bottiglia.cl
			}
		}
	}
	fmt.Println("n. bottiglie:", numBottiglie)
	fmt.Println("bottiglia con grado max:", maxGradiBottiglia)
	fmt.Println("bottiglia più vecchia:", maxAnnoBottiglia)
	fmt.Println("tot vino:", clTot, "cl")
}
