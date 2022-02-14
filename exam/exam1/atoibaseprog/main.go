package main

import (
	"fmt"
)

func CheckBase(base string) bool {
	var newstr string
	if len(base) < 2 {
		return false
	}
	for _, l := range base {
		if l == '+' || l == '-' {
			return false
		}
		for _, l2 := range newstr {
			if l == l2 {
				return false
			}
		}
		newstr += string(l)
	}
	return true
}

func len(base string) int {
	lenght := 0
	for range base {
		lenght++
	}
	return lenght
}

func AtoiBase(s string, base string) int {
	if !CheckBase(base) {
		return 0
	}
	res := 0
	for i1 := range s {
		for i2 := range base {
			if s[i1] == base[i2] {
				res = res*len(base) + i2
			}
		}
	}
	return res
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "ab"))
}
