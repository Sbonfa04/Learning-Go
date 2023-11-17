package main

func main() {
	type data struct {
		g, m, a int 
	}

	type studente struct {
		nome, cognome string
		dataDiNascita data
		matricola string
		media float64
	}

	var x studente 
	/*ogni variabile quindi bisogna immaginarsela un po' tipo database
	dove sarà contenuto il nome, cognome, dataDiNNascita 
	(a sua volta divisa in g,m,a) e il resto*/

	x.nome = "Samuele"
	x.cognome = "Bonfanti"
	x.dataDiNascita.g = 29
	x.dataDiNascita.m = 04
	x.dataDiNascita.a = 2004
	x.matricola = "A29905"
	x.media = 21.4
}