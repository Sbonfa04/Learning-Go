package main

import (
	"fmt"
)

type Persona struct {
	nome string
}

func (p *Persona) cambiaNome(n string) {
	p.nome = n
}

func (p Persona) cambiaNome2(n string) {
	p.nome = n
}

func main() {
	p := Persona{"Giovanni"}
	fmt.Println(p.nome)
	p.cambiaNome("Mario")
	fmt.Println(p.nome)
	p.cambiaNome2("Luigi")
	fmt.Println(p.nome)
}
