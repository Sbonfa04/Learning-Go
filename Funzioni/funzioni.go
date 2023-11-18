/*Per compilare tutti i file > go build fun.go fun2.go fun3.go ...
Per dare un nome all'eseguibile > go build -o "nome exe" *.go
(se ci sono solo i file da eseguire nella directoru corrente) */

package main

import (
	. "fmt"
)

func main() {

	var x float64
	Println("inserire un numero decimale: ")
	Scan(&x)

	/*per richiamare l'altra funzione*/
	separa(x)
}
