package main

import (
	"os"

	"github.com/01-edu/z01"
)

func len(input string) int {
	lengh := 0
	for range input {
		lengh++
	}
	return lengh
}

func main() {
	lenghofargs := 0
	for range os.Args[1:] {
		lenghofargs++
	}
	if lenghofargs == 0 {
		return
	}
	if lenghofargs != 2 {
		z01.PrintRune(10)
		return
	} else {
		str1 := os.Args[1]
		sep := os.Args[2]
		var str string
		str += str1 + sep
		r1 := []rune(str)
		l1 := len(string(r1))
		r2 := []rune(sep)
		l2 := len(string(r2))
		nbword := 0
		for i := 0; i <= l1-l2; i++ {
			if str[i:i+l2] == sep {
				nbword++
			}
		}
		arr := make([]string, nbword)
		counter := 0
		j := 0
		for i := counter; i <= l1-l2; i++ {
			if str[i:i+l2] == sep {
				arr[j] = str[counter:i]
				counter = i + l2
				j++
			}
		}
		z01.PrintRune('[')
		for i := range arr {
			for _, l := range arr[i] {
				z01.PrintRune(l)
			}
			if i < nbword-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
		z01.PrintRune(10)
	}
}
