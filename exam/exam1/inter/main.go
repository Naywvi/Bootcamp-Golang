package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 2 {
		z01.PrintRune(10)
		return
	} else {
		str1 := os.Args[1]
		str2 := os.Args[2]
		var result string
		for _, l1 := range str1 {
			for _, l2 := range str2 {
				flag := true
				if l1 == l2 {
					for _, l3 := range result {
						if l1 == l3 {
							flag = false
						}
					}
					if flag == true {
						result = result + string(l1)
					}
				}
			}
		}

		for _, l := range result {
			z01.PrintRune(l)
		}
		z01.PrintRune(10)
	}
}
