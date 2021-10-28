package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, char := range arg[0] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
