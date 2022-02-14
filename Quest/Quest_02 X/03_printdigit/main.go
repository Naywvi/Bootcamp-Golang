package main

import "github.com/01-edu/z01"

func main() {
	boucle := 47
	for boucle < 57 {
		boucle++
		z01.PrintRune(rune(boucle))
	}
	//TODO: <= Print les valeurs 48 à 57, soit de 0 à 9 ASCII
	z01.PrintRune('\n')
}
