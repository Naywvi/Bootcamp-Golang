package main

import (
	"os"

	"github.com/01-edu/z01"
)

func LastWord(s string) {
	var start int
	var res string
	var final int
	for i := range s {
		if i != 0 && s[i] != ' ' && s[i-1] == ' ' {
			start = i
			final = i + 1
		} else if s[i] != ' ' /*&& (i == len(s)-1 || s[i+1] == ' ')*/ {
			final = i + 1
		}
		// else if s[i] != ' ' && s[i+1] == ' ' {
		// 	final = i + 1
		// }
	}
	res = string(s[start:final])
	for _, k := range res {
		z01.PrintRune(k)
	}
}

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune(10)
		return
	}
	LastWord(os.Args[1])
	z01.PrintRune(10)
}
