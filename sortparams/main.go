package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	SLICE := []rune{}
	arg := os.Args[1:]
	str := ""
	Change := '0'
	for _, x := range arg {
		str = str + x
	}
	for _, x := range str {
		SLICE = append(SLICE, x)
	}
	for i := 0; i < len(arg); i++ {
		for j := 0; j < len(arg)-i-1; j++ {
			if SLICE[j] > SLICE[j+1] {
				Change = SLICE[j]
				SLICE[j] = SLICE[j+1]
				SLICE[j+1] = Change
			}
		}
	}
	for i := 0; i < len(SLICE); i++ {
		z01.PrintRune(SLICE[i])
		z01.PrintRune('\n')
	}
}
