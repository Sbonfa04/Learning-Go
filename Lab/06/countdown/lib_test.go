/*
	"libreria" di test per gli esami, attenzione a modificare questo file!
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"testing"
)

var HEADER string = "\n\n\n" + strings.Repeat(" ", (diffwidth-64)/2) + "___   ---   ===   ^^^   ***   TEST   ***   ^^^   ===   ---   ___"

//var HEADER string = "\n\n\n" + "___   ---   ===   ^^^   ***   TEST   ***   ^^^   ===   ---   ___"

func TestCompila(t *testing.T) {
	fmt.Println(HEADER)
	fmt.Println()
	fmt.Print("Verifico compilazione... ")

	// assumo che il go si chiami come la dir e il test sia <nomesorgente>_test.go

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("*** errore nella lettura della directory corrente ***")
		t.Fail()
		return
	}
	/*
		if strings.Contains(wd, "trent") {
			return
		}
	*/
	//fmt.Println(wd)
	nomeexe := path.Base(wd) // strippato diventa nome eseguibile
	nomego := nomeexe + ".go"
	//nometest := nomeexe + "_test.go"

	fexe, err := os.Stat(nomeexe)
	if fexe == nil {
		fmt.Println("*** c'è qualche problema sul nome della directory o del file .go (non corrispondenza con le specifiche?) ***")
		fmt.Println(err)
		t.Fail()
		return
	}

	tExe := fexe.ModTime()
	//fmt.Println(nomeexe, tExe)

	fgo, _ := os.Stat(nomego)
	tGo := fgo.ModTime()
	//fmt.Println(nomego, tGo)

	//ftest, _ := os.Stat(nometest)
	//tTest := ftest.ModTime()
	//fmt.Println(nometest, tTest)

	if tGo.After(tExe) {
		fmt.Println("**************************************************************************")
		fmt.Println("*** ATTENZIONE! il sorgente non è stato compilato dopo le modifiche!!! ***")
		fmt.Println("**************************************************************************")
		t.Fail()
	} else {
		fmt.Println("L'eseguibile è AGGIORNATO")
	}
}

/* a tendere supportare anche mac (darwin)
 */
func TestLinux(t *testing.T) {
	fmt.Println(HEADER)
	fmt.Println()
	fmt.Print("Controllo sistema operativo...", runtime.GOOS)
	if runtime.GOOS == "linux" {
		fmt.Println(" OK!")
	} else {
		fmt.Println()
		fmt.Println("*************************************************")
		fmt.Println("* ATTENZIONE! sistema operativo NON supportato! *")
		fmt.Println("*************************************************")
	}
	//fmt.Println("--------------------------------------")
}

/*
NUOVA base?
lancia stud e oracolo (il nome è hardcoded 'oracolo') e confronta

- BISOGNA lasciare nelle dir del tema il file `oracolo` eseguibile compilato dal nostro sorgente e impacchettarglielo nel tar
- si possono scrivere SIA i test classici che alcuni con questa `confronta`
- per ora vedi cancellaParole per un esempio d'uso
*/
func ConfrontaConOracolo(
	t *testing.T,
	progname string,
	filestdinput string, // se nome vuoto viene creato un contenuto a "NIENTE"
	args ...string) {

	fmt.Println(HEADER)
	fmt.Println()
	fmt.Println("[ Questo test confronta l'output studente con l'output atteso ]")
	fmt.Println()
	fmt.Println(">>> L'eseguibile da testare (", progname, ") deve essere stato compilato! <<<")
	fmt.Println()

	fileoracolo := "./oracolo"

	// fail fast se manca eseguibile studente
	if _, err := os.Stat(progname); err != nil {
		fmt.Println(">>> Manca eseguibile studente!")
		t.Fail()
		return
	}

	// fail fast se manca oracolo
	if _, err := os.Stat(fileoracolo); err != nil {
		fmt.Println(">>> Manca oracolo!")
		t.Fail()
		return
	}

	// prep i due stdin
	stdin1, filestdinputGlob, err1 := wrapStdin(filestdinput)
	stdin2, filestdinputGlob, err2 := wrapStdin(filestdinput)

	if err1 != nil || err2 != nil {
		t.Fail()
		return
	}

	//fmt.Println(stdin)

	// lancia oracolo e cattura out
	// chiama LanciaGenericaConFileInOutAtteso - NO, meglio rifarla qui così depreco le altre e questa diventa la base

	oracolo := exec.Command(fileoracolo, args...)
	oracolo.Stdin = stdin1
	oracoloout, err := oracolo.CombinedOutput()
	if err != nil {
		fmt.Printf("Attenzione! ORACOLO uscito con codice: %s\n\t>>> (non è un test fallito se si termina il programma con un esplicito os.Exit)\n", err)
	}
	//fmt.Println(">>> oracolo:", string(oracoloout))

	studente := exec.Command(progname, args...)
	studente.Stdin = stdin2
	studenteout, err := studente.CombinedOutput()
	if err != nil {
		fmt.Printf("Attenzione! STUDENTE uscito con codice: %s\n\t>>> (non è un test fallito se si termina il programma con un esplicito os.Exit)\n", err)
	}
	//fmt.Println(">>> studente:", string(studenteout))

	fmt.Printf("/// Argomenti a linea di comando:\t%s\n", args)
	fmt.Println()
	fmt.Printf("/// File per StdInput:\t%s\n", filestdinputGlob)
	fmt.Println()
	fmt.Println("### eseguo diff...")
	fmt.Println()
	e := Diff2strings(string(studenteout), "studente", string(oracoloout), "oracolo")

	//fmt.Printf("\n/// Output:\n%s\n", string(stdout))
	//fmt.Printf("\n/// Output atteso:\n%s\n", expectedOutString)

	//if string(stdout) != oracolo {
	if e != nil {
		fmt.Println(">>> FAIL! differisce da output atteso")
		t.Fail()
	}

	oracolo.Process.Kill()
	studente.Process.Kill()
	fmt.Println()
}

func wrapStdin(filestdinput string) (stdin io.Reader, filestdin string, err error) {
	filestdin = filestdinput

	if len(filestdinput) == 0 {
		stdin = strings.NewReader("NIENTE") // dummy
		filestdin = "<non era previsto input da stdin>"
	} else {
		stdin, err = os.Open(filestdinput)
		if err != nil {
			fmt.Println(">>> Non posso aprire file stdin:", filestdin)
		}
	}
	return
}

/*
TODO FATTORIZZARE LE VARIE LANCIA... ?
- in ingresso c'è stdin (potenzialmente vuoto) stringa/nomefile, args (potenzialmente vuoto)
- in output catturo stdout
- per testare confronto stdout con oracolo
- stdin posso usare direttamente i filename!!!
*/

/*
la base è tutto in forma di stringa (in e oracolo)
*/
func LanciaGenerica(
	t *testing.T,
	progname string,
	strinput string,
	oracolo string,
	args ...string) {

	//fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>> DEPRECATA? <<<<<<<<<<<<<<<<<<<<<<<<<<")

	fmt.Println(HEADER)
	fmt.Println()
	fmt.Println("[ Questo test confronta l'output con l'output atteso ]")
	fmt.Println("[ Presuppone che l'eseguibile da testare (", progname, ") sia già stato compilato! ]")
	fmt.Println()

	subproc := exec.Command(progname, args...)
	subproc.Stdin = strings.NewReader(strinput)
	stdout, err := subproc.CombinedOutput() // invece di Run()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)
		//fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
		fmt.Printf("Attenzione! Uscito con codice: %s\n\t>>> (non è un test fallito se si termina il programma con un esplicito os.Exit)\n", err)
	}

	//fmt.Println()

	fmt.Printf("\n/// Argomenti a linea di comando:\n\t%s\n", args)
	fmt.Printf("\n/// StdInput:\n%s\n", strinput)
	fmt.Println("/// eseguo diff...")
	e := Diff2strings(string(stdout), "studente", oracolo, "atteso")

	//fmt.Printf("\n/// Output:\n%s\n", string(stdout))
	//fmt.Printf("\n/// Output atteso:\n%s\n", expectedOutString)

	//if string(stdout) != oracolo {
	if e != nil {
		fmt.Println(">>> FAIL! differisce da output atteso")
		t.Fail()
	}

	subproc.Process.Kill()
	fmt.Println()
}

/*
si carica tutto in memoria... :(
*/
func LanciaGenericaConFileOutAtteso(
	t *testing.T,
	nomeProg string,
	strinput string,
	oracoloFilename string,
	args ...string) {

	content, err := ioutil.ReadFile(oracoloFilename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	//fmt.Println(text)

	LanciaGenerica(t, nomeProg, strinput, text, args...)
}

func LanciaGenericaConFileInOutAtteso(
	t *testing.T,
	nomeProg string,
	inputFilename string,
	oracoloFilename string,
	args ...string) {

	input, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	in := string(input)

	exout, err := ioutil.ReadFile(oracoloFilename)
	if err != nil {
		log.Fatal(err)
	}
	out := string(exout)
	//fmt.Println(text)

	LanciaGenerica(t, nomeProg, in, out, args...)
}

/*
vede se nell'eseguibile c'è linkata la funzione
*/
func checkIfUsed(prog, fun string) bool {
	cmd := "strings " + prog + " | grep " + fun + "|cat" // se termina con grep (che ritorna >0 se non trova) dà err!=nil
	//fmt.Println(cmd)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", cmd)
		fmt.Println(err)
		return false
	}
	//fmt.Println(string(out))
	//fmt.Println(len(out))

	if len(out) > 0 {
		return true
	}
	return false
}

func intorno(a, b float64) bool {
	return math.Abs(a-b) < 10e-6
}

func Diff2files(fn1, fn2 string) (e error) {
	//cmd := exec.Command("diff", "-y", fn1, fn2)
	cmd := exec.Command("diff", "-y", "-W", fmt.Sprint(diffwidth), fn1, fn2)
	//cmd := exec.Command("diff", "-y","-W 200", "--color=always", fn1, fn2) // verificare se c'è opzione color dappertutto

	beg := "BEGIN DIFF"
	end := " END DIFF "
	pad := ((diffwidth - len(beg)) / 2) - 1
	fmt.Println(strings.Repeat("-", pad), beg, strings.Repeat("-", pad))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e = cmd.Run()
	fmt.Println(strings.Repeat("-", pad), end, strings.Repeat("-", pad))
	return
}

/*
func OLD_Diff2files(fn1, fn2 string) {
	cmd := exec.Command("diff", "-y", fn1, fn2)
	//cmd := exec.Command("diff", "-y", "-W 200", fn1, fn2)
	//cmd := exec.Command("diff", "-y","-W 200", "--color=always", fn1, fn2) // verificare se c'è opzione color dappertutto

	fmt.Println("[ SX:", fn1, "- DX:", fn2, "]")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println()
}
*/

/*
(inefficiente, lo so) crea due file temp, ci rovescia le due stringhe e chiama altro diff
*/
func Diff2strings(str1, l1, str2, l2 string) (e error) {
	//TODO val la pena fattorizzare?
	tmpfile1, err1 := ioutil.TempFile("", l1+".*")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(tmpfile1.Name()) // clean up
	if _, err1 := tmpfile1.Write([]byte(str1)); err1 != nil {
		log.Fatal(err1)
	}
	if err1 := tmpfile1.Close(); err1 != nil {
		log.Fatal(err1)
	}

	tmpfile2, err2 := ioutil.TempFile("", l2+".*")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer os.Remove(tmpfile2.Name()) // clean up
	if _, err2 := tmpfile2.Write([]byte(str2)); err2 != nil {
		log.Fatal(err2)
	}
	if err2 := tmpfile2.Close(); err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(strings.Repeat("-", diffwidth))
	//fmt.Println("   <<< SINISTRA:", l1, "--- DESTRA:", l2, ">>>")
	fmt.Println("[", l1, "]", strings.Repeat(" ", diffwidth-len(l1)-len(l2)-10), "[", l2, "]")
	//fmt.Println(strings.Repeat("-", diffwidth))
	return Diff2files(tmpfile1.Name(), tmpfile2.Name())
}

/* non funziona su un subprocess
func consoleSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(string(out))
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return heigth, width
}
*/
