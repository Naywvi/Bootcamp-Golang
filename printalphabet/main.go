package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 97; compteur < 123; compteur++ {
		z01.PrintRune(rune(compteur))
	}
	z01.PrintRune('\n')
}
