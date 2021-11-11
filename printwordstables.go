package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	RSLICE := []rune{}
	str := ""
	for i := range a {
		str = a[i] + " "
		for _, x := range str {
			RSLICE = append(RSLICE, x)
		}
		str = ""
	}
	for i := 0; i < len(RSLICE); i++ {
		if RSLICE[i] == ' ' {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(RSLICE[i])
		}
	}
}
