/*
Vini
----

Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

  - una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
    nome string
    anno int
    gradi float32
    cl int

  - una funzione
    CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool)
    che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a questi dati e la restituisce insieme al valore 'true',
    altrimenti restituisce una bottiglia "zero-value" e 'false'.

  - una funzione
    CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
    che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo CSV (vedi esempio nelle specifiche del main);
    se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e la restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
    Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

  - un **metodo** per BottigliaVino
    String() string
    che restituisce una riga di descrizione della bottiglia nel seguente formato:  <nome>, <anno>, <gradi>°, <cl>cl
    (cioè ad es. "Rasol, 2018, 14°, 750cl" per la prima riga dell'esempio sopra).
    Suggerimento: i "format verb" %g e %v formattano i float omettendo il punto decimale quando non ci sono cifre dopo la virgola

- una funzione main() che legge da un file (il cui nome è passato da linea di comando) delle righe che contengono ognuna i dati relativi ad una bottiglia di una cantina, separati da virgole, nel formato del seguente esempio (nome,anno,gradi,cl):

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

e stampa su stdout:

  - l'elenco delle bottiglie della cantina (esattamente nello stesso formato rappresentato qui sopra).
    Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

  - il numero di bottiglie nella cantina

  - la bottiglia con gradazione massima

  - la bottiglia più vecchia

  - i cl totali di vino della cantina

Esempio di esecuzione con input vini.input
---------------------

Rasol, 2018, 14°, 750cl
Camunnorum, 2015, 15°, 750cl
Dom Perignon, 2019, 12.5°, 1500cl
Balon, 2013, 15°, 750cl
Verdicchio, 2020, 11°, 375cl
Rasol, 2018, 14°, 1000cl
Verdicchio, 2020, 11°, 375cl
n. bottiglie: 7
bottiglia con grado max: Camunnorum, 2015, 15°, 750cl
bottiglia più vecchia: Balon, 2013, 15°, 750cl
tot vino: 5500 cl
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome  string
	anno  int
	gradi float32
	cl    int
}

func CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool) {
	if nome != "" && anno > 0 && gradi > 0 && cl > 0 {
		return BottigliaVino{nome, anno, gradi, cl}, true
	} else {
		return BottigliaVino{}, false
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	splitRiga := strings.Split(riga, ",")
	if len(splitRiga) != 4 {
		return BottigliaVino{}, false
	}
	nome := splitRiga[0]
	anno, err := strconv.Atoi(strings.TrimSpace(splitRiga[1]))
	if err != nil {
		return BottigliaVino{}, false
	}
	gradi, err := strconv.ParseFloat(strings.TrimSpace(splitRiga[2]), 32)
	if err != nil {
		return BottigliaVino{}, false
	}
	cl, err := strconv.Atoi(strings.TrimSpace(splitRiga[3]))
	if err != nil {
		return BottigliaVino{}, false
	}
	return CreaBottiglia(nome, anno, float32(gradi), cl)
}

func (b BottigliaVino) String() string {
	return fmt.Sprintf("%s, %d, %g°, %dcl", b.nome, b.anno, b.gradi, b.cl)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run vini.go <file>")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var bottiglie []BottigliaVino
	var maxGrad float32
	var maxGradBott BottigliaVino
	var oldBott BottigliaVino
	var oldYear int
	var totalCl int
	for scanner.Scan() {
		b, ok := CreaBottigliaDaRiga(scanner.Text())
		if ok {
			bottiglie = append(bottiglie, b)
			if b.gradi > maxGrad {
				maxGrad = b.gradi
				maxGradBott = b
			}
			if oldYear == 0 || b.anno < oldYear {
				oldYear = b.anno
				oldBott = b
			}
			totalCl += b.cl
			fmt.Println(b)
		}
	}
	fmt.Println("n. bottiglie:", len(bottiglie))
	fmt.Println("bottiglia con grado max:", maxGradBott)
	fmt.Println("bottiglia più vecchia:", oldBott)
	fmt.Println("tot vino:", totalCl, "cl")
}
