package main

import (
	"os"

	"github.com/01-edu/z01"
)

//func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func Itoa(count int) {
	if count == 0 {
		return
	} else if count > 0 {
		Itoa(count / 10)
		z01.PrintRune(rune(count%10 + 48))
	}
}
func Atoi(str string) int {
	res := 0
	for _, l := range str {
		count := 0
		for i := '0'; i < l; i++ {
			count++
		}
		res = res*10 + count

	}
	return res
}

func Cheker(s string) bool {
	for _, l := range s {
		if l < '0' || l > '9' {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('0')
		z01.PrintRune(10)
		return
	} else {
		str := os.Args[1]
		if !Cheker(str) {
			z01.PrintRune('0')
			z01.PrintRune(10)
			return
		}
		nb := Atoi(str)
		if nb < 2 {
			z01.PrintRune('0')
			z01.PrintRune(10)
			return
		} else {
			count := 0
			for i := 2; i <= nb; i++ {
				if IsPrime(i) {
					count += i
				}
			}
			Itoa(count)
			z01.PrintRune(10)
		}
	}
}
