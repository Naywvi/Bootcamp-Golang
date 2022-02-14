package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str1 := "ddf6vewg64f"
	str2 := "gtwthgdwthdwfteewhrtag6h4ffdhsd"
	str := str1 + str2
	var res string
	for _, l1 := range str {
		flag := true
		for _, l2 := range res {
			if l1 == l2 {
				flag = false
			}
		}
		if flag == true {
			res += string(l1)
		}
	}
	for _, l3 := range res {
		z01.PrintRune(l3)
	}
	z01.PrintRune(10)
}
