package main

func integer() {
	type intero = int //diamo a "intero" il tipo intero
	var x intero
	var y int
	x = y + 5

	type inter int //int e inter sono due tipi diversi
	var a int
	var b inter
	y = int(x)
	x = inter(y)

}