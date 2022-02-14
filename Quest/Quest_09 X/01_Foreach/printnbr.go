package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune(rune(48))
	}
	if n < 0 && n != -9223372036854775808 {
		if n < 0 {
			n *= -1
		}
		SLICE := []int{}
		ns := 0

		for n != 0 {
			if n != 0 {
				ns = n % 10
				n /= 10
				SLICE = append(SLICE, ns)
			} else {
				return
			}
		}
		z01.PrintRune(rune(45))
		nr := len(SLICE)
		for i := nr - 1; i > 0; i-- {
			z01.PrintRune(rune(SLICE[i] + 48))
		}
		z01.PrintRune(rune(SLICE[0] + 48))
	} else if n > 0 {
		SLICE := []int{}
		ns := 0
		for n != 0 {
			if n != 0 {
				ns = n % 10
				n /= 10
				SLICE = append(SLICE, ns)
			} else {
				return
			}
		}
		nr := len(SLICE)
		for i := nr - 1; i > 0; i-- {
			z01.PrintRune(rune(SLICE[i] + 48))
		}
		z01.PrintRune(rune(SLICE[0] + 48))
	} else if n == -9223372036854775808 {
		SLICE := []int{}
		ns := 0

		for n != 0 {
			if n != 0 {
				ns = n % 10 * (-1)
				n /= 10
				SLICE = append(SLICE, ns)
			} else {
				return
			}
		}
		z01.PrintRune(rune(45))
		nr := len(SLICE)
		for i := nr - 1; i > 0; i-- {
			z01.PrintRune(rune(SLICE[i] + 48))
		}
		z01.PrintRune(rune(SLICE[0] + 48))
	}
}
