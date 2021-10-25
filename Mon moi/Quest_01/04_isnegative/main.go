package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {

	if nb >= 0 {
		z01.PrintRune(rune(70))
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(rune(84))
		z01.PrintRune('\n')
	}
	//TODO: <= Si nb est superieur ou égale à zero donc 'F' pour false. Sinon else et 'T' pour True
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
