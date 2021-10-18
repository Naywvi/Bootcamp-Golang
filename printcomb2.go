package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	a := 0
	b := 0
	c := 0
	d := 0

	for i := 0; i < 10000; i++ {
		d++
		if d == 10 {
			c++
			d = 0
		}
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
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
		}

		if a*10+b < c*10+d {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(32)
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		}
	}
}
