package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	count := 0
	for s := range arg {
		count = s + 1
	}
	for i := 1; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for j := 1; j <= count-1; j++ {
		for _, element := range arg[j] {
			z01.PrintRune(element)
		}
		z01.PrintRune('\n')
	}
}
