package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str1 := "fgex.;"
	str2 := "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"

	for i := range str2 {
		if i < len(str2)-len(str1) && str2[i:i+len(str1)] == str1 {
			z01.PrintRune('1')
			z01.PrintRune(10)
			return
		}
	}
	z01.PrintRune('0')
	z01.PrintRune(10)
}
