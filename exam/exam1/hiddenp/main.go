package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 2 {
		z01.PrintRune(10)
		return
	} else {
		str1 := os.Args[1]
		str2 := os.Args[2]
		var str string
		counter := 0
		for _, l1 := range str1 {
			flag := true
			for i := counter + 1; i < len(str2); i++ {
				if flag == true {
					if l1 == rune(str2[i]) {
						str += string(l1)
						counter++
						flag = false
					}
				}

			}
		}
		//fmt.Println(str)
		if str == str1 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune(10)
	}
}
