package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := []rune(os.Args[0])[2:]
	for _, x := range s {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}

// Ici nous ne voulons pas l'index 0 d'ou le 2: nous aurions pu mettre 1: Cependant Ytrack bug d'ou le 2.
