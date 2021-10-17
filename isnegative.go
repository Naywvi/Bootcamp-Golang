package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune(84)
	} else {
		z01.PrintRune(70)
	}
	z01.PrintRune('\n')
}
