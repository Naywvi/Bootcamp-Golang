package main

import "fmt"

func ConcatParams(args []string) string {
	line := "\n"
	str := ""
	for i := 0; i < len(args); i++ {
		if i != len(args)-1 {
			str += (args[i] + line)
		} else {
			str += (args[i])
		}
	}
	return str
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
