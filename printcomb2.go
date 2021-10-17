package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	a := 0
	b := 0
	c := 0
	d := 0

	for i := 0; i < 10000; i++ { //Boucle s'appelant i allant de 1 à 10k
		d++          //D prend 1
		if d == 10 { //Incrémentation > Quand d prend une dizaine c prend 1 mais d reste à zéro
			c++
			d = 0
		}
		if c == 10 { //Incrémentation > Quand c prend une dizaine b prend 1 mais c reste à zéro
			b++
			c = 0
		}
		if b == 10 { //Incrémentation > Quand b prend une dizaine a prend 1 mais b reste à zéro
			a++
			b = 0
		}
		if a == 9 && b == 8 && c == 9 && d == 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(32))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune('\n'))
			break
		} //Isolation du caractère final soit 98 99

		if a*10+b < c*10+d {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(32)
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		} //Si Ax10+b < cx10+d
	}
}

/*
func main() {
	PrintComb2()
}
*/
