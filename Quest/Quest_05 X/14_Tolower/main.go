package main

import "fmt"

func ToLower(s string) string {
	relou := []rune(s)
	for x := range relou {
		if relou[x] >= 'A' && relou[x] <= 'Z' {
			relou[x] = relou[x] + 32
		}
	}
	return string(relou)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
