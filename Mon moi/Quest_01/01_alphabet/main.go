package main

import "github.com/01-edu/z01"

func main() {
	boucle := 96
	for boucle < 122 {
		boucle++
		z01.PrintRune(rune(boucle))
	}
	z01.PrintRune('\n')
}
