package main

import (
	"fmt"
	"testing"
)

var prog = "./rettangolo"
var diffwidth = 120

/*
func TestTODO(t *testing.T) {
	fmt.Println("RIMPOLPARE TEST!!!", prog)
	t.Fail()
}
*/

func TestMainMinimale(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"NIENTE",
		"............\n............\n............\n............\n\n",
		"12", "4")
}

func TestEsistenzaStructEMetodo(t *testing.T) {
	var r Rettangolo
	//fmt.Println(r)
	fmt.Println(r.String())
}

func TestNormale(t *testing.T) {
	var r Rettangolo
	r.base = 7
	r.altezza = 5
	//fmt.Println(r)
	fmt.Println(r.String())
	if r.String() != ".......\n.......\n.......\n.......\n.......\n" {
		t.Fail()
	}
}

func TestQuadrato(t *testing.T) {
	var r Rettangolo
	r.base = 5
	r.altezza = 5
	//fmt.Println(r)
	fmt.Println(r.String())
	if r.String() != ".....\n.....\n.....\n.....\n.....\n" {
		t.Fail()
	}
}

func TestMinimo(t *testing.T) {
	var r Rettangolo
	r.base = 1
	r.altezza = 1
	//fmt.Println(r)
	fmt.Println(r.String())
	if r.String() != ".\n" {
		t.Fail()
	}
}

func TestDegenere(t *testing.T) {
	var r Rettangolo
	fmt.Println(r.String())
	if r.String() != "rettangolo degenere" {
		t.Fail()
	}
}

func TestDegenereB0(t *testing.T) {
	var r Rettangolo
	r.altezza = 8
	fmt.Println(r.String())
	if r.String() != "rettangolo degenere" {
		t.Fail()
	}
}
func TestDegenereA0(t *testing.T) {
	var r Rettangolo
	r.base = 5
	fmt.Println(r.String())
	if r.String() != "rettangolo degenere" {
		t.Fail()
	}
}
