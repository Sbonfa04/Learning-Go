/*dati due studenti restituisca il cognome dello 
studente più bravo*/

package main

type studente struct {
	nome, cognome string
	media float64
}

func student(x,y studente) string {
	if x.media > y.media {
		return x.cognome
	} else {
		return y.cognome
	}
}


func main() {
	var x1, x2 studente
	x1.nome = "Samu"
	x1.cognome = "Bonfa"
	x1.media = 30

	x2.nome = "Edo"
	x2.cognome = "Cara"
	x2.media = 29.9
}