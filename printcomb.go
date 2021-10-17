package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var a, b, c int
	for i := 0; i < 789; i++ { //Boucle s'appelant i allant de 1 à 799
		c++ //C prend 1
		if i == 788 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune('\n'))
			break
		} //Si i = 788 soit la fin on print et on saute une ligne
		if c == 10 { //Incrémentation > Quand c prend une dizaine b prend 1 mais c reste à zéro
			b++
			c = 0
		}
		if b == 10 { //Incrémentation > Quand b prend une dizaine a prend 1 mais b reste à zéro
			a++
			b = 0
		}
		if c > b && c > a && b > a {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		} //Ici on prend seulement les c>b , c>a et b>a
	}
}

/*
	func main() {
		PrintComb()
	}
*/
