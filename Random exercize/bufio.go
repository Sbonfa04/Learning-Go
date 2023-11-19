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
	for scanner.Scan() {
		riga := scanner.Text()
		if len(riga) < 1 {
			break
		}
		Println("letta riga di <", riga, ">, di lunghezza", strconv.Itoa(len(riga)), "bytes")
	}
}
