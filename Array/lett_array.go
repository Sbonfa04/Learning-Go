package main

func main() {
	x := [8]int{1, 2, 3, 4, 5}                           //array lungo 8 con 5 valori dentro
	y := [...]int{6, 7, 8, 9, 10}                        //array lungo quanto la quantità di elementi contenuti (in questo caso 5)
	z := [40]string{1: "ciao", 7: "pippo", 39: "ultimo"} //inserisco solo alcuni elementi in posizioni specifiche
}
