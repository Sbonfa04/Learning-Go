package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	ascissa   float64
	ordinata  float64
}

type duePunti struct {
	uno Punto
	due Punto
}

type Triangolo struct {
	A Punto
	B Punto
	C Punto
}

func LeggiPunti() (punti []Punto) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), ";")
		ascissa, _ := strconv.ParseFloat(splitted[1], 64)
		ordinata, _ := strconv.ParseFloat(splitted[2], 64)

		punti = append(punti, Punto{splitted[0], ascissa, ordinata})
	}

	return
}

func GeneraTriangoli(punti []Punto) (triangoli []Triangolo) {

	permutazioniDuePunti := make([]duePunti, 0)
	//PERMUTAZIONI A 2 PUNTI

	for i := 0; i < len(punti); i++ {
		for j := i + 1; j < len(punti); j++ {
			permutazioniDuePunti = append(permutazioniDuePunti, duePunti{punti[i], punti[j]})
		}
	}

	//PERMUTAZIONI A 3 PUNTI

	for _, v := range punti {
		for _, f := range permutazioniDuePunti {
			present := false
			if v != f.uno && v != f.due { //se il punto non è già uguale a uno dei due nella coppia

				for _, triangolo := range triangoli {
					//se non c'è già un triangolo composto dagli stessi punti (anche in ordine diverso)
					if (triangolo.A == v || triangolo.B == v || triangolo.C == v) && (triangolo.A == f.uno || triangolo.B == f.uno || triangolo.C == f.uno) && (triangolo.A == f.due || triangolo.B == f.due || triangolo.C == f.due) {
						present = true
						break
					}
				}
				if !present {
					triangoli = append(triangoli, Triangolo{v, f.uno, f.due})
				}
			}
		}
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.ascissa-p2.ascissa, 2) + math.Pow(p1.ordinata-p2.ordinata, 2))
}

func areaTriangolo(t Triangolo) float64 {
	l1 := Distanza(t.A, t.B)
	l2 := Distanza(t.A, t.C)
	l3 := Distanza(t.B, t.C)

	perimetro := l1 + l2 + l3
	s := perimetro / 2

	area := math.Sqrt(s * (s - l1) * (s - l2) * (s - l3))
	return area
}

func StringPunto(p Punto) string {
	return fmt.Sprintf("%s = (%f,%f)", p.etichetta, p.ascissa, p.ordinata)
}

func StringTriangolo(t Triangolo) string {
	return fmt.Sprintf("Triangolo rettangolo con vertici %s,%s,%s, ed area %f", StringPunto(t.A), StringPunto(t.B), StringPunto(t.C), areaTriangolo(t))
}

func FiltraTriangoli(triangoli []Triangolo) (triangoliFiltrati []Triangolo) {
	triangoliFiltrati = make([]Triangolo, 0)
	for _, triangolo := range triangoli {
		//per ogni triangolo
		//se almeno un lato ha x uguali e almeno un lato ha y uguali (è rettangolo)
		if (triangolo.A.ascissa == triangolo.B.ascissa || triangolo.A.ascissa == triangolo.C.ascissa || triangolo.B.ascissa == triangolo.C.ascissa) && (triangolo.A.ordinata == triangolo.B.ordinata || triangolo.A.ordinata == triangolo.C.ordinata || triangolo.B.ordinata == triangolo.C.ordinata) {
			//è rettangolo, controllare se anche tutti i punti sono nello stesso quadrante, cioè e y o le x hanno stesso segno.
			if (math.Signbit(triangolo.A.ascissa) == math.Signbit(triangolo.B.ascissa) && math.Signbit(triangolo.B.ascissa) == math.Signbit(triangolo.C.ascissa)) && (math.Signbit(triangolo.A.ordinata) == math.Signbit(triangolo.B.ordinata) && math.Signbit(triangolo.B.ordinata) == math.Signbit(triangolo.C.ordinata)) {
				triangoliFiltrati = append(triangoliFiltrati, triangolo) //se tutte le x hanno stesso segno, e anche tutte le y hanno stesso segno, allora il triangolo è tutto in un quadrante.
			}
		}
	}
	return
}

func main() {
	/*1.
	uno dei due cateti del triangolo rettangolo sia parallelo all'asse delle ascisse (l'altro cateto
	deve essere quindi parallelo all'asse delle ordinate); UN LATO DEVE AVERE LE X UGUALI E UN LATO DEVE AVERE LE Y UGUALI
	2.i vertici del triangolo rettangolo giacciano tutti nello stesso quadrante del piano cartesiano; LE X DEVONO AVERE STESSO SEGNO O LE Y DEVONO AVERE STESSO SEGNO
	3.abbia area maggiore tra tutti i triangoli rettangoli che soddisfano le condizioni 1 e 2. FILTRARE TRIANGOLI

	Si ricorda che è possibile calcolare l'area di un triangolo rettangolo come: area = cateto1
	cateto2 / 2 oppure con la formula di Erone: area = (p (p - l1) (p - l2) (p - l3))1/2
	*/

	triangoliRettangoliStessoQuadrante := FiltraTriangoli(GeneraTriangoli(LeggiPunti()))

	maxArea := 0.0
	maxTriangolo := Triangolo{}

	if len(triangoliRettangoliStessoQuadrante) != 0 {
		for _, triangolo := range triangoliRettangoliStessoQuadrante {

			area := areaTriangolo(triangolo)

			if area > maxArea {
				maxArea = area
				maxTriangolo = triangolo
			}

		}

		fmt.Println(StringTriangolo(maxTriangolo))
	}
}
