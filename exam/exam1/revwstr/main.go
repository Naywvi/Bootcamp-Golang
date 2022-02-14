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
		str := os.Args[1]
		var newstr string
		var start = 0
		var final int
		for i1 := 0; i1 < len(str)-1; i1++ {
			if i1 < len(str)-1 && str[i1] != ' ' && str[i1+1] == ' ' {
				final = i1 + 1
				newstr = str[start:final] + " " + newstr
				start = final + 1
			}
		}
		newstr = str[start:] + " " + newstr
		newstr = newstr[0 : len(newstr)-1]
		for _, l := range newstr {
			z01.PrintRune(l)
		}
		z01.PrintRune(10)
	}
}
