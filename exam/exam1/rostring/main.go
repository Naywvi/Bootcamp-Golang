package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str := "Let there     be light  "
	var newstr string
	var firstword string
	for i, l := range str {
		if l != ' ' {
			newstr += string(l)
		} else if i != 0 && str[i-1] != ' ' && l == ' ' {
			newstr += string(l)
		}
	}
	flag := true
	count := 0
	for _, l1 := range newstr {
		if l1 == ' ' {
			flag = false
		}
		if flag == true {
			firstword += string(l1)
			count++
		}
	}
	for i := count + 1; i < len(newstr); i++ {
		z01.PrintRune(rune(newstr[i]))
	}
	if str[len(str)-1] != ' ' {
		z01.PrintRune(' ')
	}
	for _, l := range firstword {
		z01.PrintRune(l)
	}
	z01.PrintRune(10)
}
