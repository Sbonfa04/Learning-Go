/*
Scrivere un programma carrello.go che simula gli spostamenti di un carrello lungo un percorso lineare a celle unitarie, in cui sono collocati degli oggetti di peso specificato.
Il carrello deve avanzare di una cella alla volta e ripulire il percorso caricando gli oggetti man mano che li incontra, ma senza superare il suo carico massimo.
Il carrello, quando si trova su una cella in cui c'è un oggetto, deve caricarlo; se non può caricarlo (il controllo lo fa sempre e solo mentre è sulla cella precedente), non deve avanzare su quella cella.
Quando l'oggetto sulla cella successiva farebbe superare il carico massimo, il carrello, invece di procedere, deve andare all'inizio del percorso, scaricare lì (nella cella 0) tutto il suo carico, e ripartire a pulire il percorso dagli oggetti che ancora vi si trovano.
Una volta caricato l'ultimo oggetto e ripulito tutto il percorso, il carrello deve svuotare il suo carico nella cella 0 e portarsi sulla cella 1.

Il programma deve essere dotato di:

- una struttura Carrello con campi
	- caricoMax int
	- posizione int
	- carico int
  dichiarati in quest'ordine

- un metodo String() string
  per Carrello che restituisce una descrizione del carrello che specifica posizione, carico e carico massimo, nel formato:
  posizione <posizione>, carico <carico> (max <caricoMax>)
  Ad esempio, per un Carrello con carico massimo 50 kg, che si trova sulla cella 25 e sta trasportando 42 kg, restituisce la stringa:
  "posizione 25, carico 42 (max 50)"

- una funzione aggiornaStato(c *Carrello, posizione, carico int) bool
  che aggiorna i dati del carrello (numero di cella e carico) e restuisce true se posizione e carico non sono negativi e se il carico non supera il carico massimo; altrimenti non fa nessun aggiornamento e restituisce false.

- una funzione main() che legge da file, il cui nome è passato su linea di comando, la descrizione di un percorso (a celle) con oggetti da rimuovere (vedi esempio sotto).
Una cella è delimitata da '|' e le celle sono numerate a partire da 0.
Se una cella se è vuota (se c'è uno spazio bianco tra '|'), in quella posizione non c'è nessun oggetto; se invece la cella contiene un numero, in quella posizione c'è un oggetto di quel peso.

Ad esempio, la seguente è la descrizione di un percorso:

"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"

che ha un oggetto di peso 12 nella cella 3, uno di peso 4 nella cella 4, ecc. come mostrato qui sotto:

"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"
  0 1 2 3  4 5 6 7  8 9 10 .....

La funzione main deve istanziare un Carrello che porta un massimo di 15 kg, si trova nella cella 0 del percorso, e ha un carico pari a 0, e poi gestire il Carrello in modo che ripulisca il percorso letto da file, operando come descritto sopra.
Si può dare per scontato, senza fare controlli, che sul percorso non ci siano pesi maggiori del peso massimo che il carrello può caricare.

Il main deve stampare:
- il percorso come è prima di iniziare il lavoro di pulizia;
- il carrello ogni volta che si verifica il caso che non è possibile caricare un nuovo oggetto;
- il percorso dopo ogni scarico;
- il carrello alla fine del lavoro;
- il numero di viaggi indietro fatti per rimuovere tutti gli oggetti dal percorso;
- il massimo peso trovato lungo il percorso;
- per ogni peso (in ordine crescente di peso), il numero di oggetti trovati di quel peso.

(Nota. si consiglia in fase di sviluppo di stampare anche la situazione del carrello dopo ogni carico)

Nel caso manchi il nome del file, il programma deve stampare il messaggio "manca nome file" e terminare.

Esempio di esecuzione per il percorso dato sopra
------------------------------------------------

| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 3, carico 12 (max 15)
|12| | | |4| | | |10| | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 12, carico 14 (max 15)
|26| | | | | | | | | | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 19, carico 9 (max 15)
|35| | | | | | | | | | | | | | | | | | | |12| | | | |3| |
|50| | | | | | | | | | | | | | | | | | | | | | | | | | |
carrello: posizione 1, carico 0 (max 15)
n viaggi: 4
peso max: 12
oggetti trovati:
1 ogg. di peso 3
2 ogg. di peso 4
1 ogg. di peso 5
1 ogg. di peso 10
2 ogg. di peso 12



*****************************************************************
>>> NOTA BENE

I test qui sotto funzionano correttamente su Linux e con
l'oracolo fornito, ma dato che l'oracolo è compilato per Linux
tutti i test basati su di esso non funzioneranno (falliranno)
su altre piattaforme.
Ergo sono stati forniti anche dei test basati sui
file ".expected" per chi usa Windows o Mac a casa.
*****************************************************************

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Carrello struct {
	caricoMax int
	pos       int
	carico    int
}

func (c Carrello) String() string {
	return fmt.Sprintf("posizione %v, carico %v (max %v)", c.pos, c.carico, c.caricoMax)
}

func aggiornaStato(c *Carrello, posizione, carico int) bool {
	if posizione < 0 || carico < 0 || carico > c.caricoMax {
		return false
	}
	c.pos = posizione
	c.carico = carico
	return true
}

func stampaPercorso(path []int, c Carrello) string {
	var strPath []string
	for _, num := range path {
		str := strconv.Itoa(num)
		if num != 0 {
			strPath = append(strPath, str)
		} else {
			strPath = append(strPath, " ")
		}
	}
	strPath = strPath[1 : len(strPath)-len(strPath)/2]
	return ("|" + strconv.Itoa(c.carico) + "|" + strings.Join(strPath[1:], "|") + "|")
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var c Carrello
	var maxPeso, nViaggi int
	var oggetti [100]int
	//cereate a slice that contains the path taken in the file where each element is one cell of the path
	//the path is the first line of the file
	//the path is a slice of int where each element if it's 0 there is no object in that cell, otherwise there is an object of that weight
	var strPath []string
	var path []int
	//scan the first line of the file and create the path
	scanner.Scan()
	strPath = strings.Split(scanner.Text(), "|")
	//convert the path from string to int
	for _, str := range strPath {
		num, _ := strconv.Atoi(str)
		path = append(path, num)
	}
	path = path[1 : len(path)-1]

	for _, num := range path {
		str := strconv.Itoa(num)
		if num != 0 {
			strPath = append(strPath, str)
		} else {
			strPath = append(strPath, " ")
		}
	}
	strPath = strPath[1 : len(strPath)-len(strPath)/2]
	// between every element and at the start and the finish of strPath put a "|" and print it as a string
	fmt.Println("|" + strings.Join(strPath, "|") + "|")
	//create a carrello with max weight 15, position and weight decided in the file (the cart is the first cell of the path that contains an object)
	for i := 0; i < len(path); i++ {
		if path[i] != 0 {
			c = Carrello{15, i, path[i]}
			break
		}
	}
	//print the carrello status
	fmt.Println("carrello:", c)

	//for each cell of the path
	for i := 0; i < len(path); i++ {
		//if the cell contains an object
		if path[i] != 0 {
			//check if the cart can load the object
			if c.carico+path[i] <= c.caricoMax {
				//if it can load the object, update the cart status with the new position and weight
				aggiornaStato(&c, i, c.carico+path[i])
				//remove the object from the string path
				strPath[i+1] = " "
				//if the cart is in the last cell of the path
				if i == len(path)-1 {
					//update the cart status with the new position and weight
					aggiornaStato(&c, 1, 0)
					//print the new path with the cart in the first position with the sum of the weight of the objects loaded and with the object removed
					fmt.Println(stampaPercorso(path, c))
					//print the cart status
					fmt.Println("carrello:", c)
					//update the max weight
					if c.carico > maxPeso {
						maxPeso = c.carico
					}
				}
			} else {
				//if it can't load the object, update the cart status with the new position and weight
				aggiornaStato(&c, 0, c.carico)
				//remove the old cart from the string path
				strPath[c.pos+1] = " "
				//print the updated path with the cart in the first position with the sum of the weight of the objects loaded and with the object removed
				fmt.Println(stampaPercorso(path, c))
				//print the cart status
				fmt.Println("carrello:", c)
				//update the max weight
				if c.carico > maxPeso {
					maxPeso = c.carico
				}
				//update the number of trips
				nViaggi++
			}
		} else {
			//if the cell doesn't contain an object, update the cart status with the new position and weight
			aggiornaStato(&c, i, c.carico)
		}
	}
	//print the number of trips
	fmt.Println("n viaggi:", nViaggi)
	//print the max weight
	fmt.Println("peso max:", maxPeso)
	//count the number of objects for each weight
	for i := 0; i < len(path); i++ {
		if path[i] != 0 {
			oggetti[path[i]]++
		}
	}
	//print the number of objects for each weight
	fmt.Println("oggetti trovati:")
	for i := 0; i < len(oggetti); i++ {
		if oggetti[i] != 0 {
			fmt.Println(oggetti[i], "ogg. di peso", i)
		}
	}
}
