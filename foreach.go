package piscine

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		PrintNbr(a[i])
	}
	z01.PrintRune('\n')
}
