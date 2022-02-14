package main

import (
	"github.com/01-edu/z01"
)

func IsPrime(nb int) bool {
	if nb == 0 || nb == 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 && nb != i {
			return false
		}
	}
	return true
}
func addPrimeSum(nb int) int {
	res := 0
	for i := 2; i <= nb; i++ {
		if IsPrime(i) {
			res += i
		}
	}
	return res
}

func Itoa(nb int) {
	if nb == 0 {
		return
	} else if nb > 0 {
		Itoa(nb / 10)
		z01.PrintRune(rune(nb%10 + 48))
	}
}

func Atoi(str string) int {
	res := 0
	for _, l := range str {
		counter := 0
		for i := '0'; i <= l; i++ {
			counter++
		}
		res = res*10 + counter
	}
	return res
}

func main() {
	nb := "7"
	Itoa(addPrimeSum(Atoi(nb)))
	z01.PrintRune(10)
}
