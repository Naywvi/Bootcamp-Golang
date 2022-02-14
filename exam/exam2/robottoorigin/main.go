package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
	rest := "true"
	resf := "false"
	var ox int //Up(+1) and Down(-1)
	var oy int // Left(-1) and Right(+1)
	for _, letter := range args {
		if letter == 'U' {
			ox++
		}
		if letter == 'D' {
			ox--
		}
		if letter == 'R' {
			oy++
		}
		if letter == 'L' {
			oy--
		}
	}
	if ox == 0 && oy == 0 {
		for _, l := range rest {
			z01.PrintRune(l)
		}
		z01.PrintRune(10)
	} else {
		for _, l := range resf {
			z01.PrintRune(l)
		}
		z01.PrintRune(10)
	}
}
