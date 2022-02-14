package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1]
	for _, l := range arr {
		if l >= 'A' && l <= 'M' || l >= 'a' && l <= 'm' {
			l = l + 13
			z01.PrintRune(l)
		} else if l >= 'N' && l <= 'Z' || l >= 'n' && l <= 'z' {
			l = l - 13
			z01.PrintRune(l)
		} else {
			z01.PrintRune(l)
		}
	}
	z01.PrintRune(10)
}
