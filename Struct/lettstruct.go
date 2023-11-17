package main

type data struct {
	g, m, a int
}


func boldi() {
	var d data
	var s studente

	d = data{29, 11, 1968}
	s = studente{
		"Paolo",
		"Boldi",
		data{29, 11, 1968},
		"352372",
		18.6
	}
}
