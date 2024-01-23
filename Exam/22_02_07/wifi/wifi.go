/*

Wifi
----

Si scriva un programma (il file deve chiamarsi 'wifi.go') per gestire l'analisi dei segnali wifi di un ambiente.
Il programma dovrà essere dotato delle seguenti:

- una struttura Wifi con campi (dichiarati in quest'ordine):
	ssid      	string
	channel   	int
	frequency 	int
	signal_dBm  int

- una funzione NewWifi(ssid string, channel int, frequency int, signal_dBm int) Wifi,bool
	che, se i valori passati come parametri rispettano le seguenti condizioni:
	- ssid non vuoto
	- channel
		- tra 1 e 14 (compresi) SE la frequency è tra 2412 e 2484 (compresi)
		- OPPURE tra 7 e 196 SE la frequency è tra 5035 e 5980
	- frequency tra 2412 e 2484 OPPURE tra 5035 e 5980 (compresi estremi)
	- signal_dBm minore di -20 (meno venti)
	crea un'istanza di Wifi opportunamente inizializzata e la restituisce insieme a true, altrimenti restituisce una struttura "zero-value" e false

- una funzione NewWifiDaStringa(line string) Wifi,bool
	che costruisce un'istanza della struttura Wifi a partire da una stringa nel formato CSV (in cui i dati
	sono separati da virgole, vedere il file 'wifi.csv'), stessi vincoli della funzione definita sopra.
	Il formato è esattamente come nel file, non occorre fare controlli sul formato, ma i dati potrebbero
	essere non accettabili (ad es. un numero di canale non coerente con la frequenza o l'intestazione del CSV).

- un metodo String da applicare a Wifi
	che restituisca una stringa rappresentativa dei valori della struttura, nella forma:
		{ssid, channel, frequency, signal_dBm, signalW}
	Attenzione che c'è un valore in più rispetto ai dati della struct, va aggiunto un valore calcolato: la potenza del segnale in Watt. Il formato di uscita del valore signalW è quello "naturale" del float64 (formato "%v").

- una funzione ConvertiDBinWatt(signal_dBm int) float64
	che prende come parametro la potenza del segnale in dBm (decibel-milliwatts) e calcola la potenza in Watt, la formula di conversione è:
		Watt = (10^(potenza_in_dBm/10)) / 1000
		Nota: il simbolo ^ indica elevamento a potenza (10^2 è 10 alla seconda)

- una funzione PiuPotente(elenco []Wifi, banda string) int
	che restituisce l'indice della struct che rappresenta il wifi con  segnale (dBm) più potente
	dell'elenco passato come parametro; in funzione del parametro banda viene restituito il più potente
	fra i segnali a 2GHz (banda="2GHz"), fra quelli a 5GHz (banda="5GHz") o senza distinzione
	(banda = qualunque altro valore, compresa la stringa vuota)

- una funzione main()
	che elabora un file (CSV) il cui nome è passato come primo argomento (della linea di comando) e che
	stampa i dati del segnale più potente sulla base del secondo argomento: se NON presente o se diverso
	da "5GHz"/"2GHz", trova il più potente senza distinzione di banda; se invece il secondo argomento è
	presente, stampa i dati secondo la banda richiesta.
	Non occorre fare controlli sulla presenza degli argomenti sulla linea di comando e sul file, potete
	assumere che il programma venga sempre lanciato correttamente e che il file indicato sia presente e
	nel formato previsto (vedi sopra).

Esempi di esecuzione
---------------------

$ ./wifi wifi.csv
{at-wind,11,2462,-41,7.94328234724282e-08}

$ ./wifi wifi.csv 5GHz
{at-wind,108,5540,-42,6.30957344480193e-08}

$ ./wifi wifi.csv 2GHz
{at-wind,11,2462,-41,7.94328234724282e-08}

*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Wifi struct {
	ssid       string
	channel    int
	frequency  int
	signal_dBm int
}

func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi, bool) {
	if ssid == "" {
		if frequency >= 2412 && frequency <= 2484 {
			if channel >= 1 && channel <= 14 {
				if signal_dBm < -20 {
					return Wifi{ssid, channel, frequency, signal_dBm}, true
				} else {
					return Wifi{}, false
				}
			} else {
				return Wifi{}, false
			}
		} else if frequency >= 5035 && frequency <= 5980 {
			if channel >= 7 && channel <= 196 {
				if signal_dBm < -20 {
					return Wifi{ssid, channel, frequency, signal_dBm}, true
				} else {
					return Wifi{}, false
				}
			} else {
				return Wifi{}, false
			}
		} else {
			return Wifi{}, false
		}
	} else {
		return Wifi{}, false
	}
}

func NewWifiDaStringa(line string) (Wifi, bool) {
	var ssid string
	var channel int
	var frequency int
	var signal_dBm int

	strings := strings.Split(line, ",")
	ssid = strings[0]
	channel, _ = strconv.Atoi(strings[1])
	frequency, _ = strconv.Atoi(strings[2])
	signal_dBm, _ = strconv.Atoi(strings[3])

	wifi, flag := NewWifi(ssid, channel, frequency, signal_dBm)

	return wifi, flag
}

func (w Wifi) String() string {
	return fmt.Sprintf("{%s,%d,%d,%d,%v}", w.ssid, w.channel, w.frequency, w.signal_dBm, ConvertiDBinWatt(w.signal_dBm))
}

func ConvertiDBinWatt(signal_dBm int) float64 {
	return math.Pow(10, float64(signal_dBm)/10) / 1000
}

func PiuPotente(elenco []Wifi, banda string) int {
	if banda == "2GHz" {
		var max float64 = 0
		var indice int = 0
		for i, v := range elenco {
			if v.frequency >= 2412 && v.frequency <= 2484 {
				if v.signal_dBm > int(max) {
					max = ConvertiDBinWatt(v.signal_dBm)
					indice = i
				}
			}
		}
		return indice
	} else if banda == "5GHz" {
		var max float64 = 0
		var indice int = 0
		for i, v := range elenco {
			if v.frequency >= 5035 && v.frequency <= 5980 {
				if v.signal_dBm > int(max) {
					max = ConvertiDBinWatt(v.signal_dBm)
					indice = i
				}
			}
		}
		return indice
	} else {
		var max float64 = 0
		var indice int = 0
		for i, v := range elenco {
			if v.signal_dBm > int(max) {
				max = ConvertiDBinWatt(v.signal_dBm)
				indice = i
			}
		}
		return indice
	}
}

func main() {
	file, _ := os.Open(os.Args[1])

	var elenco []Wifi
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		riga := scanner.Text()
		wifi, flag := NewWifiDaStringa(riga)
		if flag {
			elenco = append(elenco, wifi)
		}
	}

	if len(os.Args) == 2 {
		if os.Args[2] == "2GHz" {
			fmt.Println(elenco[PiuPotente(elenco, "2GHz")])
		} else if os.Args[2] == "5GHz" {
			fmt.Println(elenco[PiuPotente(elenco, "5GHz")])
		} else {
			fmt.Println(elenco[PiuPotente(elenco, "")])
		}
	}
}
