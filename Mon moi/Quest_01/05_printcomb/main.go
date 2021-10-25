package main

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	a := 0
	b := 1
	c := 2
	boucle := 0
	for boucle < 1000 {
		boucle++
		c++
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
			a++
			b = 0
		}
		//TODO: <= Incrémentation
		if a == 7 && b == 8 && c == 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(44))
			z01.PrintRune(rune(32))
			z01.PrintRune('\n')
			break
		}
		//TODO: <= Cas particulier, si 789 = print = fin (break pas oblighatoire)
		if a < b && b < c {

			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(44))
			z01.PrintRune(rune(32))
		}
		//TODO: <= Condition pour print a<b<c (car 0<1<2)
	}

}

func main() {
	PrintComb()
}
