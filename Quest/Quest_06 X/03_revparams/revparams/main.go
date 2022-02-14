package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i >= 1; i-- {
		for _, x := range arg[i] {
			z01.PrintRune(x)
		}
		z01.PrintRune('\n')
	}
}
