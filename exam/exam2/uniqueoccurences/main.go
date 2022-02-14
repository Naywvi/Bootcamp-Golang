package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	m := make(map[rune]int)
	for _, l := range Inter(args) {
		m[l] = countOccur(l, args)
	}
	for k1, v1 := range m {
		for k2, v2 := range m {
			if k1 != k2 && v1 == v2 {
				fmt.Println(false)
				return
			}
		}
	}
	fmt.Println(true)
}

func Inter(str string) string {
	var res string
	for _, l1 := range str {
		flag := true
		for _, l2 := range res {
			if l1 == l2 {
				flag = false
			}
		}
		if flag == true {
			res += string(l1)
		}
	}
	//fmt.Println(res)
	return res
}

func countOccur(r rune, args string) int {
	var counter int
	for _, l := range args {
		if l == r {
			counter++
		}
	}
	return counter
}
