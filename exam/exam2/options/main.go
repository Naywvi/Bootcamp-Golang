// 0	"os.*"
// 1	"fmt.*"
// 2	"len"
// 3	"--cast"
package main

import (
	"fmt"
	"os"
)

func CheckStr(str string) bool {
	for _, l := range str {
		if l < 'a' || l > 'z' {
			return false
		}
	}
	return true
}

func CheckHelp(args []string) bool {
	for _, str := range args {
		if str[0] == '-' && str[1] == 'h' {
			return true
		}
	}
	return false
}

func PrintRes(res [32]int) {
	counter := 0
	for i := 31; i >= 0; i-- {
		fmt.Print(res[i])
		counter++
		if i == 0 {
			fmt.Println()
			return
		}
		if counter == 8 {
			fmt.Print(" ")
			counter = 0
		}
	}
}

func main() {
	var res [32]int
	if CheckHelp(os.Args[1:]) {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	for _, str := range os.Args[1:] {
		if CheckStr(str[1:]) && str[0] == '-' {
			for _, l := range str[1:] {
				res[int(l)-97] = 1
			}
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}
	PrintRes(res)
}
