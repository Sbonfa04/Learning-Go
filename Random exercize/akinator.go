package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	question string
	yes, no  *Node
}

func initialize() *Node {
	root := new(Node)
	yesAnswer := new(Node)
	noAnswer := new(Node)
	*root = Node{"E' una ragazza?", yesAnswer, noAnswer}
	*yesAnswer = Node{"Giorgia Meloni", nil, nil}
	*noAnswer = Node{"Matteo Salvini", nil, nil}
	return root
}

/* Legge da standard input una risposta (yes or no) e restituisce o 'y' o 'n' */
func yes_no() rune {
	var answer string
	for {
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		if len(answer) == 0 || answer[0] != 'y' && answer[0] != 'n' {
			fmt.Printf("\tPlease answer y(es) or n(o)... ")
		} else {
			return rune(answer[0])
		}
	}
}

func ask(root *Node) *Node {
	curr := root
	for curr.yes != nil {
		fmt.Printf("%s [y/n] ", curr.question)
		yn := yes_no()
		if yn == 'y' {
			curr = curr.yes
		} else {
			curr = curr.no
		}
	}
	return curr
}

func play(root *Node) {
	scanner := bufio.NewScanner(os.Stdin)

	leaf := ask(root)
	fmt.Printf("Per me stai pensando a %s\n", leaf.question)
	fmt.Printf("Ho ragione? ")
	yn := yes_no()
	if yn == 'y' {
		fmt.Println("Suka ho vinto :)")
	} else {
		fmt.Printf("A chi stavi pensando? ")
		scanner.Scan()
		rightAnswer := scanner.Text()
		fmt.Printf("Dimmi una domanda in cui la risposta per %s sia sì e per %s sia no: ", leaf.question, rightAnswer)
		scanner.Scan()
		question := scanner.Text()

		nodeYes := new(Node)
		nodeNo := new(Node)
		*nodeYes = Node{leaf.question, nil, nil}
		*nodeNo = Node{rightAnswer, nil, nil}
		*leaf = Node{question, nodeYes, nodeNo}
	}
}

func main() {
	var root *Node

	root = initialize()
	for {
		play(root)
		fmt.Println()
	}
}
