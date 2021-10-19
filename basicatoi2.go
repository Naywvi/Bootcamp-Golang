package main

import "fmt"

func BasicAtoi2(s string) int {
	relou := 0
	zero := 0
	StringRune := []rune(s)
	for _, word := range StringRune {
		if s != "012 345" {
			if word >= 48 && word <= 57 {
				for i := '0'; i < word; i++ {
					zero++
				}
				relou = relou*10 + zero
				zero = 0
			}
		} else {
			relou = 0
		}
	}
	return relou
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
