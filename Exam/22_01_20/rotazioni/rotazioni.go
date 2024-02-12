/*
Rotazioni
========

Si scriva un programma (il file deve chiamarsi 'rotazioni.go') per gestire rotazioni di rettangoli nel piano cartesiano, come spiegato qui di seguito.

Il programma dovrà essere dotato di:

- una struttura Point con campi x, y di tipo float64 (dichiarati in quest'ordine)

  - una struttura Rectangle con campi pLL, pUR di tipo Point (dichiarati in quest'ordine),
    dove pLL è il verice in basso a sinistra (LowerLeft) e pUR quello in alto a destra (Upper Right) del rettangolo

  - una funzione NewPoint(x, y float64) Point
    che data una coppia di coordinate x,y, restituisce il punto in quella posizione

  - una funzione NewRectangle(p1, p2 Point) Rectangle
    che dati due vertici opposti di un rettangolo (P1 e P3 o P2 e P4, indifferentemente),
    restituisce il rettangolo corrispondente (vedi definizione di Rectangle), cioè in cui i due campi
    sono il vertice in basso a sinistra e quello il alto a destra (indicati con P1 e P3 nella figura sotto)

Nota: indipendentemente dai valori delle coordinate di due vertici opposti di un rettangolo

	(P1 e P3 o P2 e P4), P1 è dato da (min(x1,y1), min(y1,y2) e P3 da (max(x1,x1), max(y1,y2)

	             P4 _____________ P3
	               |             |
	               |             |
	               |_____________|
	             P1               P2

- una funzione Rotate(r *Rectangle, dir byte)
che ruota il rettangolo r intorno al suo vertice pLL di 90 gradi in senso orario se dir è 'R' e in senso antiorario se dir è 'L'.
Per esempio, ruotando il rettangolo qui sopra in direzione 'L', il rettangolo diventa quello della figura
qui sotto (P1 del vecchio rettangolo corrisponde a P2 del nuovo).
Per direzioni diverse da 'L' e 'R' la funzione non fa niente.

		P4 _______ P3
		  |       |
		  |       |
		  |       |
		  |       |
		  |_______|
		P1         P2

	  - un ***metodo***  da applicare a Rectangle
	    String() string
	    che restituisce una riga di descrizione del rettangolo nel seguente formato: ((P1.x,P1.y) (P3.x,P3.y))
	    (cioè ad es. "((2.1,3) (7,13.5))"
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Rectangle struct {
	pLL, pUR Point
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func NewRectangle(p1, p2 Point) Rectangle {
	return Rectangle{p1, p2}
}

func distance(p1, p2 Point) (float64, float64) {
	return math.Abs(p1.x - p2.x), math.Abs(p1.y - p2.y)
}

func Rotate(r *Rectangle, dir byte) {
	if dir == 'R' {
		dx, dy := distance(r.pLL, r.pUR)
		r.pUR.x = r.pLL.x + dy
		r.pUR.y = r.pLL.y
		r.pLL.y = r.pLL.y - dx
	} else if dir == 'L' {
		dx, dy := distance(r.pLL, r.pUR)
		r.pUR.x = r.pLL.x
		r.pUR.y = r.pLL.y + dx
		r.pLL.x = r.pLL.x - dy
	}
}

func (r Rectangle) String() string {
	return fmt.Sprintf("((%.6g,%.6g) (%.6g,%.6g))", r.pLL.x, r.pLL.y, r.pUR.x, r.pUR.y)
}

func main() {

}
