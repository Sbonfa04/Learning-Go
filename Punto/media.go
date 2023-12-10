package main

func Media(p1, p2 Punto) Punto {
	var p Punto
	p.X, p.Y = (p1.X+p2.X)/2, (p1.Y+p2.Y)/2
	return p
}
