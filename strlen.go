package piscine

import "github.com/01-edu/z01"

func StrLen(s string) {
	var w int = 0
	for _, s := range s {
		w++
		if s == 1 {
			w++
			s = 12
		}
		if w == 12 {
			z01.PrintRune(49 + 50)
		}

	}
}
