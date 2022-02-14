package piscine

import "fmt"

func BasicAtoi2(s string) int {
	relou := 0
	StringRune := []rune(s)
	for i := 0; i < len(s); i++ {
		if StringRune[i] < '0' || StringRune[i] > '9' {
			return 0
		} else {
			relou *= 10
			relou += int(StringRune[i]) - '0'
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
