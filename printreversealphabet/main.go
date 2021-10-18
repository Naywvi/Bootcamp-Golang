package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 122; compteur > 96; compteur-- {
		z01.PrintRune(rune(compteur))
	}
	z01.PrintRune(rune('\n'))
}
