package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	t := "true"
	f := "false"
	var result bool
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	} else {
		nb, err := strconv.Atoi(os.Args[1])
		if err != nil {
			str1 := "strconv.Atoi: parsing \""
			str2 := "\": invalid syntax\n"
			str := str1 + os.Args[1] + str2
			for _, i := range str {
				z01.PrintRune(i)
			}
			return
		} else {
			if nb == 0 {
				for _, l2 := range f {
					z01.PrintRune(l2)
				}
				z01.PrintRune(10)
				return
			}
			if nb == 1 {
				result = true
			}
			for i := nb; i > 1; i = i / 2 {
				if i%2 == 1 || i == 0 {
					for _, l2 := range f {
						z01.PrintRune(l2)
					}
					z01.PrintRune(10)
					return
				} else {
					result = true
				}
			}
			if result == true {
				for _, l1 := range t {
					z01.PrintRune(l1)
				}
				z01.PrintRune(10)
				return
			}
		}
	}
}
