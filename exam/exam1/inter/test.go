package main

import (
	"github.com/01-edu/z01"
)

func main() {
	str1 := "padinton"
	str2 := "paqefwtdjetyiytjneytjoeyjnejeyj"

	for _, l1 := range str1 {
		flag := false
		for _, l2 := range str2 {
			if l1 == l2 {
				flag = true
			}
		}
		if flag == true {
			z01.PrintRune(l1)
		}
	}
	z01.PrintRune(10)
}
