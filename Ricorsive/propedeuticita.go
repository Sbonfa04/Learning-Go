package main

type Esami struct {
	nome string
	t    []string
}

func main() {
	var esame1, esame2, esame3 Esami
	esame1 = Esami{
		"1",
		[]string{""},
	}
	esame2 = Esami{
		"2",
		[]string{""},
	}
	esame3 = Esami{
		"3",
		[]string{"1"},
	}

	var arrayEsami []Esami
	arrayEsami = append(arrayEsami, esame3, esame2, esame1)
}

func ordineEsami(e []Esami) []string {
	var output []string

	if e == []Esami{} {
		return []string{""}
	} else {
		for i := 0; i < len(output); i++ {
			if e[0].t[j] == output[i] {
				output
			}
		}

	}
}
