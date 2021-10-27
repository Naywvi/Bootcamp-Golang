package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune(rune(48))
	}
	if n > 0 {
		if n > 0 {
			for boucle2 := n; boucle2 > 0; boucle2-- {
				for boucle := 0; boucle < 3; boucle++ {
					r := n / 10
					x := n % 10
					n = r
					z01.PrintRune(rune(48 + x))
				}
				if n == 0 {
					z01.PrintRune('\n')
					return
				}
			}
		}
	}
}
