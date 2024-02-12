package main

import (
	"fmt"
	"os"
)

type Segmento struct {
	estremi     byte
	interno     byte
	orizzontale bool
	lunghezza   int
}

func (s Segmento) String() string {
	var str string
	if s.orizzontale {
		if s.lunghezza >= 2 {
			str = string(s.estremi)
			for j := 0; j < s.lunghezza-2; j++ {
				str += string(s.interno)
			}
			str += string(s.estremi)
		} else {
			if s.lunghezza == 1 {
				str = string(s.estremi)
			}
		}
	} else {
		if s.lunghezza >= 2 {
			str = string(s.estremi) + "\n"
			for j := 0; j < s.lunghezza-2; j++ {
				str += string(s.interno) + "\n"
			}
			str += string(s.estremi)
		} else {
			if s.lunghezza == 1 {
				str = string(s.estremi)
			}
		}
	}
	return str
}

func (s *Segmento) allunga() string {
	s.lunghezza += 3
	return s.String()
}

func main() {
	var s Segmento
	var esterni byte
	var interni byte
	var orizzontale bool
	var lunghezza int

	esterni = os.Args[1][0]
	interni = os.Args[2][0]
	o := os.Args[3]
	lunghezza = int(os.Args[4][0] - '0')

	if o == "true" {
		orizzontale = true
	} else {
		orizzontale = false
	}

	s = Segmento{esterni, interni, orizzontale, lunghezza}
	fmt.Println(s.String())
	fmt.Println(s.allunga())
	orizzontale = !orizzontale
	s = Segmento{esterni, interni, orizzontale, lunghezza}
	fmt.Println(s.String())
}
