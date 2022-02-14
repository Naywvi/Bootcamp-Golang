package main

import "fmt"

func BasicJoin(elems []string) string {
	wrld := ""
	for i := 0; i < len(elems); i++ {
		wrld = wrld + elems[i]
	}
	return wrld
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
