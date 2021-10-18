package piscine

import "github.com/01-edu/z01"

func StrLen() {
	var o string = "Hello World!"
	var l int = 0
	for _, s := range o {
		l++
		if s == 1 {
			l++
			s = 12
		}
		if l == 12 {
			z01.PrintRune(49 + 50)
		}

	}
}
