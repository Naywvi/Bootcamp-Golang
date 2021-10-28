package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, x := range arg {
		z01.PrintRune(x)
	}
}
