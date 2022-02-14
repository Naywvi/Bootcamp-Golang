package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func Itoa(nb int) {
	if nb == 0 {
		return
	} else if nb > 0 {
		Itoa(nb / 10)
		z01.PrintRune(rune(nb%10 + 48))
	}
}

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	} else {
		nb, err := strconv.Atoi(os.Args[1])
		if err != nil {
			z01.PrintRune(10)
			return
		} else if nb > 1 && nb < 2147483647 {
			for i := 1; i <= 9; i++ {
				res := i * nb
				if res < 2147483647 {
					Itoa(i)
					z01.PrintRune(' ')
					z01.PrintRune('x')
					z01.PrintRune(' ')
					Itoa(nb)
					z01.PrintRune(' ')
					z01.PrintRune('=')
					z01.PrintRune(' ')
					Itoa(res)
					z01.PrintRune(10)
				}
			}
		}
	}
}
