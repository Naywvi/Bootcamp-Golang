package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPowerof2(nb int) bool {
	if nb == 0 {
		return false
	} else if nb == 1 {
		return true
	} else {
		for i := nb; i > 1; i = i / 2 {
			if i%2 == 1 || i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	input := os.Args[1]
	nb, _ := strconv.Atoi(input)
	if isPowerof2(nb) {
		z01.PrintRune('T')
	} else if !isPowerof2(nb) {
		z01.PrintRune('F')
	}
	z01.PrintRune(10)
}
