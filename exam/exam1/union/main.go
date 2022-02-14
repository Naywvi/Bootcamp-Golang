package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println()
		return
	} else {
		str1 := os.Args[1]
		str2 := os.Args[2]
		str := str1 + str2
		var slice []rune
		for _, l1 := range str {
			flag := true
			for _, l2 := range slice {
				if l1 == l2 {
					flag = false
				}
			}
			if flag == true {
				slice = append(slice, l1)
			}
		}
		for _, l := range slice {
			fmt.Print(string(l))
		}
		fmt.Println()
	}
}
