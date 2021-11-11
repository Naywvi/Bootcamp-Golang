package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	count := 0
	for i := range a {
		count++
		r := string(a[i])
		for o, x := range r {
			if count == 1 && o < len(r)-1 {
				z01.PrintRune(x)
			} else if o == len(r)-1 {
				z01.PrintRune(x)
				z01.PrintRune('\n')
				count = 0
			}
		}
	}
}
