/*data una funzione studente, ne cambia il nome in mario*/
package main



func Mario(s *studente) {
	(*s).nome = "Mario"
}

func main() {
	var x studente
	x.nome = "Pippo"
	x.cognome = "Bello"
	x.media = 22

	Mario(&x)
}