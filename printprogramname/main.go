package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := []rune(os.Args[0])
	for _, x := range s {
		z01.PrintRune(x)
	}
}
