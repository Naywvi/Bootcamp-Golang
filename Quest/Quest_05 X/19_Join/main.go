package main

import (
	"fmt"
)

func Join(strs []string, sep string) string {
	s := ""
	count := 0
	for i := 0; i < len(strs); i++ {
		count++
		if i != len(strs)-1 {
			s = s + strs[i] + sep
		} else {
			s = s + strs[i]
		}

	}
	return s
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
