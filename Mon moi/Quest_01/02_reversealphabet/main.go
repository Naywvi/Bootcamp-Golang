package main

import "github.com/01-edu/z01"

func main() {
	boucle := 122
	for boucle > 97 {
		boucle--
		z01.PrintRune(rune(boucle))
	}
	//TODO: <= Boucle-- car la boucle va de 122 à 97
	z01.PrintRune('\n')
}
