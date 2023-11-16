package main

import (
	"fmt"
)

func main() {
	var s string
	var p *string
	s = "pippo"
	p = &s
	fmt.Println(s)       //stampa pippo
	fmt.Println(p)       //stampa indirizzo di memoria di s
	fmt.Println(*p)      // = s
	fmt.Println(len(*p)) // = a lunghezza di s
	*p = "ciao"          //riassegnamo "ciao" ad s

	var a int
	var b *int
	var c **int

	a = 15
	b = &a
	c = &b
	(**c)++
	fmt.Println(a)
	/*b contiene l'indirizzo di a e c contiene l'indirizzo di b quindi
	**c = *b = a */
}
