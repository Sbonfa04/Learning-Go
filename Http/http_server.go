package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"
	"os"

	"github.com/holizz/terrapin"
)

func pippoPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<!doctype html>\n<title>Pippo</title>\n<h1>Pippo</h1>\n<p>Oh bella Pippo</p1>\n"))
}

// disegnare su html
func turtleGraphics(w http.ResponseWriter, r *http.Request) {
	image := image.NewRGBA(image.Rect(0, 0, 500, 500))
	t := terrapin.NewTerrapin(image, terrapin.Position{180.0, 340.0})
	t.Forward(80)
	t.Right(math.Pi / 2)
	t.Forward(160)
	t.Left(math.Pi / 2)
	t.Forward(80)

	p := terrapin.NewTerrapin(image, terrapin.Position{340.0, 340.0})
	p.Left(math.Pi / 2)
	p.Forward(80)
	p.Right(math.Pi / 2)
	p.Forward(160)
	p.Left(math.Pi / 2)
	p.Forward(80)

	png.Encode(w, image)
}

func pipoImage(w http.ResponseWriter, r *http.Request) {
	image := image.NewRGBA(image.Rect(0, 0, 500, 500))

	for i := 0.0; i < (math.Pi * 2); i += math.Pi / 2000 {
		p2 := terrapin.NewTerrapin(image, terrapin.Position{180.0, 420.0})
		p2.Right(i)
		p2.Forward(70)

		p1 := terrapin.NewTerrapin(image, terrapin.Position{320.0, 420.0})
		p1.Right(i)
		p1.Forward(70)
	}

	for k := 0.0; k < (math.Pi * 2); k += math.Pi / 2000 {
		c := terrapin.NewTerrapin(image, terrapin.Position{250.0, 100.0})
		c.Right(k)
		c.Forward(50)
	}

	for j := 0; j <= 80; j++ {
		f := terrapin.NewTerrapin(image, terrapin.Position{210.0, 420.0})
		f.Forward(320)
		f.Right(math.Pi / 2)
		f.Forward(float64(j))
		f.Right(math.Pi / 2)
		f.Forward(320)
		f.Right(math.Pi / 2)
		f.Forward(float64(j))
	}

	png.Encode(w, image)

}

// inserire immagine saricata su pc
func plutoPage(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("")
	image, _ := png.Decode(f)
	png.Encode(w, image)
}

func main() {
	http.HandleFunc("/pipo", pipoImage)
	http.HandleFunc("/turtle", turtleGraphics)
	http.HandleFunc("/pippo", pippoPage)
	http.HandleFunc("/pluto", plutoPage)
	fmt.Println("Listening on http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
