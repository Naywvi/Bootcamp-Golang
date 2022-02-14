package main

import (
	"os"

	"github.com/01-edu/z01"
)

func RepeatAlpha(r rune) {
	if r >= 'A' && r <= 'Z' {
		for i := 0; i < int(r-64); i++ {
			z01.PrintRune(r)
		}
	} else if r >= 'a' && r <= 'z' {
		for i := 0; i < int(r-96); i++ {
			z01.PrintRune(r)
		}
	} else {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	} else {
		str := os.Args[1]
		for _, l1 := range str {
			RepeatAlpha(l1)
		}
	}
	z01.PrintRune(10)
}
