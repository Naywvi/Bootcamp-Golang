package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func PrintComb2() {
	boucle := 0
	a := 0
	b := 0
	c := 0
	d := 1
	for boucle < 10000 {
		boucle++
		if a == 9 && b == 8 && c == 9 && d == 9 {
			z01.PrintRune(rune(32))
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune('\n')
			break
		}
		// TODO: <= Ceci est la fin
		if b == 10 {
			a++
			b = 0
		}
		if c == 10 {
			d++
			c = 0
		}
		if d == 10 {
			c++
			d = 0
		}
		//TODO: <= Incrémentation
		if a*10+b < c*10+d {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(44))
			z01.PrintRune(rune(32))
		}
		//TODO: <= EXEMPLE : [a,b] [c,d] soit [98][99] donc [ax10+b=98|cx10+d=90] = 98 99. Une manière de rassembler deux valeurs pour en faire qu'une
	}
}

func main() {
	piscine.PrintComb2()
}
