package main

func separa(x float64) (pi int, pf float64) {
	pi = int(x)
	pf = x - float64(pi)
	return /*pi, pf
	Se non avessi dichiarato le variabili all'inizio
	dovevo speificare cosa mettere nel return*/
}
