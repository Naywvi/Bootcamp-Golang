package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str := "hello"
	slice1 := []rune(str)
	slice2 := []rune(str)
	j := len(slice2) - 1
	for i := range slice1 {
		slice2[j] = slice1[i]
		j--
	}
	for _, l := range slice2 {
		z01.PrintRune(l)
	}
	z01.PrintRune(10)
}
