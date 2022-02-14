package main

import (
	"fmt"
)

func ToUpper(s string) string {
	relou := []rune(s)
	for x := range relou {
		if relou[x] >= 'a' && relou[x] <= 'z' {
			relou[x] = relou[x] - 32
		}
	}
	return string(relou)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
