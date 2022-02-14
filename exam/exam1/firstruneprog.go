package main

import (
	"github.com/01-edu/z01"
)

func FirstRune(s string) rune {
	// var result rune
	// for i, l := range s {
	// 	if i == 0 {
	// 		result = rune(l)
	// 	}
	// }
	// return result
	len := 0
	for range s {
		len++
	}
	return rune(s[len-1])
}

func main() {
	s := "hello"
	arr := FirstRune(s)
	z01.PrintRune(arr)
	z01.PrintRune(10)
}
