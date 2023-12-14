/*

Scrivere un programma `repInPlace.go` che definisca una funzione:
	func repInPlace(stringa []rune, old rune, new rune)
che modifichi la `stringa` passata sostituendo ogni occorrenza della runa `old` con la runa `new`

Aggiungere inoltre un main che accetti tre parametri a linea di comando, rispettivamente:
- stringa da modificare
- carattere sostituendo
- carattere di rimpiazzo

e stampi la stringa letta da linea di comando e, a capo, la stessa stringa dopo che è stata modificata.

*/

package main

import (
	"fmt"
	"testing"
)

var prog = "./repInPlace"
var diffwidth = 120

/*
	func TestTODO(t *testing.T) {
		fmt.Println("RIMPOLPARE TEST!!!", prog)
		t.Fail()
	}
*/
func TestMainMinimale(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"NIENTE",
		"The Quick Brown Fox Jumps Over The Crazy Dog\nThe Quick Brown fox Jumps Over The Crazy Dog\n",
		"The Quick Brown Fox Jumps Over The Crazy Dog", "F", "f")
}

func TestMinimo(t *testing.T) {
	s := []rune("ciao")
	fmt.Println(string(s))
	repInPlace(s, 'c', 'C')
	fmt.Println(string(s))
	if s[0] != 'C' {
		t.Fail()
	}
}

func TestDoppio(t *testing.T) {
	s := []rune("ciaopippo")
	fmt.Println(string(s))
	repInPlace(s, 'c', 'C')
	repInPlace(s, 'p', 'Q')
	fmt.Println(string(s))
	if s[0] != 'C' {
		t.Fail()
	}
	if s[4] != 'Q' || s[6] != 'Q' || s[7] != 'Q' {
		t.Fail()
	}
}

func TestNormale(t *testing.T) {
	s := []rune("The Quick Brown Fox Jumps Over The Crazy Dog")
	fmt.Println(string(s))
	repInPlace(s, 'Q', 'q')
	repInPlace(s, 'B', 'b')
	repInPlace(s, 'x', 'X')
	fmt.Println(string(s), string(s[4]), string(s[10]), string(s[18]))
	if s[4] != 'q' || s[10] != 'b' || s[18] != 'X' {
		t.Fail()
	}
}
