package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 48; compteur < 58; compteur++ { //Boucle s'appelant compteur allant de 58 à 48
		z01.PrintRune(rune(compteur)) //Print en ASCII
	}
	z01.PrintRune(rune('\n'))
}
