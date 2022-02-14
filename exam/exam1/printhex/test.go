package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func Hex(nb int) rune {
	r := 'a'
	for i := 10; i < 16; i++ {
		if nb == i {
			return r
		}
		r++
	}
	return r
}

func main() {
	nb := 5156454
	var res string
	for i := nb; i > 0; i = i / 16 {
		if i%16 > 9 {
			res += string(Hex(i % 16))
		} else {
			res += strconv.Itoa(i % 16)
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(rune(res[i]))
	}
	z01.PrintRune(10)
}
