package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 122; compteur > 96; compteur-- { //Boucle s'appelant compteur allant de 122 à 96
		z01.PrintRune(rune(compteur)) //Print en ASCII
	}
	z01.PrintRune(rune('\n'))
}
