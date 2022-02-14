package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", ""))
}

func len(str string) int {
	counter := 0
	for range str {
		counter++
	}
	return counter
}

func CheckBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	var newstr string
	for _, l1 := range base {
		if l1 == '+' || l1 == '-' {
			return false
		}
		for _, l2 := range newstr {
			if l1 == l2 {
				return false
			}
		}
		newstr += string(l1)
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !CheckBase(base) {
		return 0
	} else {
		var res int
		for _, l1 := range s {
			counter := 0
			for _, l2 := range base {
				if l1 == l2 {
					break
				}
				counter++
			}
			res = res*(len(base)) + counter
		}
		return res
	}
}
