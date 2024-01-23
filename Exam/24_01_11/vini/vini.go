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
	if nome == "" || anno <= 0 || gradi <= 0 || cl <= 0 {
		return BottigliaVino{}, false
	}
	return BottigliaVino{nome, anno, gradi, cl}, true
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	var b BottigliaVino
	var nome string
	var anno, cl int
	var gradi float32
	_, err := fmt.Sscanf(riga, "%v,%d,%f,%d", &nome, &anno, &gradi, &cl)
	if err != nil {
		return b, false
	}
	return CreaBottiglia(nome, anno, gradi, cl)

}

func (b BottigliaVino) String() string {
	return fmt.Sprintf("%s, %d, %.1f°, %dcl", b.nome, b.anno, b.gradi, b.cl)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var vini []BottigliaVino
	var maxGradi BottigliaVino
	var vecchio BottigliaVino
	var totCl int

	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			continue
		}
		vino, ok := CreaBottigliaDaRiga(riga)
		if !ok {
			continue
		}
		vini = append(vini, vino)
		totCl += vino.cl
		if vino.gradi > maxGradi.gradi {
			maxGradi = vino
		}
		if vecchio.anno == 0 || vino.anno < vecchio.anno {
			vecchio = vino
		}
	}

	for _, vino := range vini {
		fmt.Println(vino)
	}
	fmt.Printf("n. bottiglie: %d\n", len(vini))
	fmt.Printf("bottiglia con grado max: %s\n", maxGradi)
	fmt.Printf("bottiglia più vecchia: %s\n", vecchio)
	fmt.Printf("tot vino: %d cl\n", totCl)
}
