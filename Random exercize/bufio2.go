package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

func main() {

	/*Lo scanner legge da una sorgente (tastiera di base) e restituisce in token
	scanner.Scan serve per richiamare se c'è la presenza di token
	se è presente scanner.Text richiama effettivamente il token */

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) /*serve per modificare il tipo di token*/
	/*in questo caso stiamo leggendo parola per parola al posto che riga per riga*/
	for scanner.Scan() {
		riga := scanner.Text()
		Println("letta riga di <", riga, ">, di lunghezza", strconv.Itoa(len(riga)), "bytes")
	}

}
