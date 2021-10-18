package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, convert := range s {
		z01.PrintRune(convert)
	}
}
