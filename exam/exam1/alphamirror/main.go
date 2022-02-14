package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	} else {
		for _, l := range os.Args[1] {
			if l == 'A' || l == 'a' {
				l = l + 25
				z01.PrintRune(l)
			} else if l == 'B' || l == 'b' {
				l = l + 23
				z01.PrintRune(l)
			} else if l == 'C' || l == 'c' {
				l = l + 21
				z01.PrintRune(l)
			} else if l == 'D' || l == 'd' {
				l = l + 19
				z01.PrintRune(l)
			} else if l == 'E' || l == 'e' {
				l = l + 17
				z01.PrintRune(l)
			} else if l == 'F' || l == 'f' {
				l = l + 15
				z01.PrintRune(l)
			} else if l == 'G' || l == 'g' {
				l = l + 13
				z01.PrintRune(l)
			} else if l == 'H' || l == 'h' {
				l = l + 11
				z01.PrintRune(l)
			} else if l == 'I' || l == 'i' {
				l = l + 9
				z01.PrintRune(l)
			} else if l == 'J' || l == 'j' {
				l = l + 7
				z01.PrintRune(l)
			} else if l == 'K' || l == 'k' {
				l = l + 5
				z01.PrintRune(l)
			} else if l == 'L' || l == 'l' {
				l = l + 3
				z01.PrintRune(l)
			} else if l == 'M' || l == 'm' {
				l = l + 1
				z01.PrintRune(l)
			} else if l == 'N' || l == 'n' {
				l = l - 1
				z01.PrintRune(l)
			} else if l == 'O' || l == 'o' {
				l = l - 3
				z01.PrintRune(l)
			} else if l == 'P' || l == 'p' {
				l = l - 5
				z01.PrintRune(l)
			} else if l == 'Q' || l == 'q' {
				l = l - 7
				z01.PrintRune(l)
			} else if l == 'R' || l == 'r' {
				l = l - 9
				z01.PrintRune(l)
			} else if l == 'S' || l == 's' {
				l = l - 11
				z01.PrintRune(l)
			} else if l == 'T' || l == 't' {
				l = l - 13
				z01.PrintRune(l)
			} else if l == 'U' || l == 'u' {
				l = l - 15
				z01.PrintRune(l)
			} else if l == 'V' || l == 'v' {
				l = l - 17
				z01.PrintRune(l)
			} else if l == 'W' || l == 'w' {
				l = l - 19
				z01.PrintRune(l)
			} else if l == 'X' || l == 'x' {
				l = l - 21
				z01.PrintRune(l)
			} else if l == 'Y' || l == 'y' {
				l = l - 23
				z01.PrintRune(l)
			} else if l == 'Z' || l == 'z' {
				l = l - 25
				z01.PrintRune(l)
			} else {
				z01.PrintRune(l)
			}
		}
	}
	z01.PrintRune(10)
}
