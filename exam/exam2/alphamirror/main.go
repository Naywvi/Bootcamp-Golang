package main

import (
	"github.com/01-edu/z01"
)

func main() {
	args := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXY"
	var counter int
	for _, letter := range args {
		if letter >= 'A' && letter <= 'Z' {
			counter = 65 + 25
		} else if letter >= 'a' && letter <= 'z' {
			counter = 97 + 25
		}
		start := counter - 25
		fin := counter
		for i := start; i <= fin; i++ {
			if letter == rune(i) {
				z01.PrintRune(rune(counter))
			}
			counter--
		}
	}
	z01.PrintRune(10)
}
