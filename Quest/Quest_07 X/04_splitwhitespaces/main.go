package main

import (
	"fmt"
)

func SplitWhiteSpaces(s string) []string {
	var lsdemots []string
	mot := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' && i != len(s)-1 {
			mot = mot + string(s[i])
		} else {
			if i == len(s)-1 {
				mot = mot + string(s[i])
			}
			if mot != "" {
				lsdemots = append(lsdemots, mot)
			}
			mot = ""
		}
	}
	return lsdemots

}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
