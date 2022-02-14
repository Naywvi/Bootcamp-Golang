package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintInt(nb int) {
	if nb == 0 {
		return
	} else if nb > 0 {
		PrintInt(nb / 10)
		z01.PrintRune(rune(nb%10 + 48))
	}
}

func StrLen(s string) {
	len := 0
	for range s {
		len++
	}
	PrintInt(len)
}
func main() {
	StrLen(os.Args[1])
	z01.PrintRune(10)
}
