package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Utenza struct {
	numeroTelefono string
	codiceSim      string
}

type RegistroTelefonico map[string][]string

func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		numero_codice := strings.Split(scanner.Text(), ";")
		utenze = append(utenze, Utenza{numero_codice[0], numero_codice[1]})
	}

	return
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {

	registroAggiornato = registro

	_, check := registroAggiornato[utenza.numeroTelefono]

	if !check {
		registroAggiornato[utenza.numeroTelefono] = make([]string, 1) //se non è presente, creo lo slice prima di aggiungere un elemento.
	}

	registroAggiornato[utenza.numeroTelefono] = append(registroAggiornato[utenza.numeroTelefono], utenza.codiceSim)

	return
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {

	_, check := registro[telefono]

	if check {
		//se c'è
		codiceSIM = registro[telefono][len(registro[telefono])-1] //ultimo elemento dello slice, se len = 1, unico elemento in pos 0.
	} else {
		codiceSIM = ""
	}

	return

}

func main() {

	elencoUtenze := LeggiUtenze()
	registro := InizializzaRegistro()

	for _, v := range elencoUtenze {
		registro = AggiungiUtenza(registro, v)
	}

	for numeroTelefono, _ := range registro {
		if numeroTelefono[:3] == "338" {
			fmt.Println("Il numero ", numeroTelefono, " è associato alla sim ", Identifica(registro, numeroTelefono))
		}
	}
}
