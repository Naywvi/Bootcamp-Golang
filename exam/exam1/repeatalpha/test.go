package main

import (
	"os"

	"github.com/01-edu/z01"
)

func RepeatAlpha(letter rune) {
	if letter >= 'A' && letter <= 'Z' {
		for i := 65; i <= 65+25; i++ {
			z01.PrintRune(letter)
			if letter == rune(i) {
				return
			}
		}
	} else if letter >= 'a' && letter <= 'z' {
		for i := 97; i <= 97+25; i++ {
			z01.PrintRune(letter)
			if letter == rune(i) {
				return
			}
		}
	} else {
		z01.PrintRune(letter)
	}
}

func main() {
	str := os.Args[1]

	for _, l := range str {
		RepeatAlpha(l)
	}
	z01.PrintRune(10)
}
