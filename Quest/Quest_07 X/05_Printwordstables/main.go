package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	count := 0
	for i := range a {
		count++
		r := string(a[i])
		for o, x := range r {
			if count == 1 && o < len(r)-1 {
				z01.PrintRune(x)
			} else if o == len(r)-1 {
				z01.PrintRune(x)
				z01.PrintRune('\n')
				count = 0
			}
		}
	}
}

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}

/*
ICI L'ERREUR A ETE DE SEPARRE CHAQUE MOT PAR UN ESPACE TANDIS QU'ON PEUT AVOIR UN ESPACE DANS LE MOT
RSLICE := []rune{}
	str := ""
	for i := range a {
		str = a[i] + " "
		for _, x := range str {
			RSLICE = append(RSLICE, x)
		}
		str = ""
	}
for i := 0; i < len(RSLICE)-1; i++ {
		if RSLICE[i] == ' ' {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(RSLICE[i])
		}
	}
*/
