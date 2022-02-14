package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func IntToRune(nb int) string {
	r := 'a'
	for i := 10; i < 16; i++ {
		if i == nb {
			return string(r)
		}
		r++
	}
	return string(r)
}

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	} else {
		input := os.Args[1]
		nb, err := strconv.Atoi(input)
		if err != nil || input == "0" {
			z01.PrintRune('0')
			z01.PrintRune(10)
			return
		} else {
			var res string
			for i := nb; i > 0; i = i / 16 {
				if i%16 < 10 {
					res += strconv.Itoa(i % 16)
				} else if i%16 > 9 {

					res += IntToRune(i % 16)
				}

			}
			for i2 := len(res) - 1; i2 >= 0; i2-- {
				z01.PrintRune(rune(res[i2]))
			}
			z01.PrintRune(10)
		}
	}
}
