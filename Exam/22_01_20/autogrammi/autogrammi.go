/*

Autogrammi
----------

Un *autogramma* è una frase che in qualche modo descrive (a volte mentendo) se stessa,
ad esempio questi che seguono sono autogrammi onesti:
"questa frase contiene cinque parole"
oppure
"questa frase contiene: 10 parole, lunghezza massima: 9, lunghezza minima: 5, esclusi numeri"
o ancora
"questa lunghissima frase invece contiene: 14 parole, lunghezza massima: 11,
lunghezza minima: 2, esclusi numeri si intende".

La frase "questa frase contiene tre parole" invece ovviamente è un autogramma bugiardo.

Dovete scrivere un programma (il cui file deve chiamarsi 'autogrammi.go')
per verificare se frasi nella forma specificata qui sotto sono autogrammi onesti o meno.

Le frasi avranno tutte le seguenti componenti, nell'ordine:
- un preambolo arbitrario ma che termina sempre con la stringa "contiene:"
- l'indicazione del numero di parole della frase nella forma "contiene: <numero> "
- l'informazione sulla lunghezza della parola più lunga nella forma "massima: <numero>,"
- l'informazione sulla lunghezza della parola più corta nella forma "minima: <numero>,"
- e eventuale altro testo di lunghezza arbitraria

Si noti anche che:
- l'unico divisore di parola è lo spazio ' ' singolo
- la punteggiatura (caratteri possibili solo ',' e ':') è sempre attaccata alla fine
delle parole, mai da sola

In particolare il programma dovrà definire le seguenti funzioni:

- CalcQuanteMinMax(frase string) (quante, min, max int)
	- conta le stringhe (separate da white space) presenti nella frase, escludendo i numeri (sequenze di sole cifre e eventuale punteggiatura attaccata alla fine)
	- calcola la lunghezza della parola più corta nella frase, **non considerando** l'eventuale punteggiatura come parte delle parole da misurare ("minima:" sarà considerata di lunghezza 6 e non 7), ed escludendo i numeri nella frase
	- calcola la lunghezza della parola più lunga nella frase, **non considerando** la punteggiatura come parte delle parole da misurare, ed escludendo i numeri nella frase
	- nota bene: per 'lunghezza' si intende il numero di byte
	- nota bene: in caso di stringa vuota, la funzione deve restituire (0,0,0)

- func TrovaNumDopoKeyword(frase, keyWord string) (num int)
	estrae da una frase il valore numerico indicato dopo una keyword e lo restituisce.
	Ad esempio, con la frase "le parole in questa frase in italiano sono == 8" e keyWord "==", restituisce 8; con la frase "questa contiene: 9" e keyword "contiene:", restituisce 9.
	Se la keyword non è presente, la funzione restituisce 0.

- func Autogramma(frase string) bool
	esamina la frase e verifica che essa sia un autogramma onesto o no; se lo è, restituisce true, false altrimenti

- func StampaAnalisiAutogramma(frase string)
	esamina la frase e stampa l'analisi effettuata secondo il formato:

		===  <frase>
		minimo dichiarato <min_dichiarato> contro minimo verificato <min_verificato>
		massimo dichiarato <max_dichiarato> contro massimo verificato <max_verificato>
		numero parole dichiarato <num_parole_dichiarato> contro numero parole verificato <num_parole_verificato>
		<onesto|bugiardo>

	senza indentazione, si veda anche esempio sottostante


- func main()
	 legge un file di testo il cui nome viene specificato a linea di comando;
	 il file contiene frasi (una per riga) che **potrebbero** essere autogrammi onesti;
	 bisogna analizzare ogni frase e stampare in output la frase letta, le informazioni ottenute dall'analisi, e "onesto" o "bugiardo" in corrispondenza delle analisi effettuate.
	 Le righe vuote vanno ignorate.

	 Se il nome del file non viene fornito, il programma stampa il messaggio "file name?" e termina.
	 Se l'apertura del file dà errore, il programma stampa il messaggio "file not found" e termina.


Esempio di esecuzione
---------------------

Se il file in input contiene le frasi:

questa frase contiene: 13 parole, lunghezza massima: 8, lunghezza minima: 5, esclusi numeri
questa più corta contiene: 11 parole, lunghezza massima: 9, lunghezza minima: 4, esclusi numeri

il programma stamperà:

===  questa frase contiene: 13 parole, lunghezza massima: 8, lunghezza minima: 5, esclusi numeri
minimo dichiarato 5 contro minimo verificato 5
massimo dichiarato 8 contro massimo verificato 9
numero parole dichiarato 13 contro numero parole verificato 10
bugiardo
===  questa più corta contiene: 11 parole, lunghezza massima: 9, lunghezza minima: 4, esclusi numeri
minimo dichiarato 4 contro minimo verificato 4
massimo dichiarato 9 contro massimo verificato 9
numero parole dichiarato 11 contro numero parole verificato 11
onesto

*/

package main

import (
	"strings"
)

func CalcQuanteMinMax(frase string) (quante, min, max int) {
	var frase []string
	frase = strings.Split(frase, " ")
}

func TrovaNumDopoKeyword(frase, keyWord string) (num int) {

}

func Autogramma(frase string) bool {

}

func StampaAnalisiAutogramma(frase string) {

}

func main() {

}
