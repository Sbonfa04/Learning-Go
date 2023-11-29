/*esegui l'anagramma di una parola con una funzione ricorsvia*/

package main

func anag(s string) []string {
	if s == "" {
		return []string{""}
	} else {
		var ris []string
		for i := 0; i < len(s); i++ {
			primo := rune(s[i])
			resto := s[:i] + s[i+1:]
			anagResto := anag(resto)

			for _, x := range anagResto {
				ris = append(ris, string(primo)+x)
			}
		}
		return ris
	}
}

func main() {}
