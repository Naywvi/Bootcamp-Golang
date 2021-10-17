package main

import "github.com/01-edu/z01"

func main() {
	for compteur := 97; compteur < 123; compteur++ { //Boucle appelé compteur allant de 97 à &é"
		z01.PrintRune(rune(compteur)) //Impréssion de la boucle en ASCII
	}
	z01.PrintRune(rune('\n'))
}
