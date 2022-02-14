package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func hiddenp(str1, str2 string) bool {
	counter1 := 0
	counter2 := 0
	for _, l1 := range str1 {
		for i := counter1; i < len(str2); i++ {
			if rune(str2[i]) == l1 {
				counter2++
				counter1 = i + 1
				break
			}
		}
	}
	fmt.Println(counter2)
	fmt.Println(len(str1))
	if len(str1) == counter2 {
		return true
	} else {
		return false
	}
}

func main() {
	str1 := "abc"
	str2 := "2altrb53c.sse"
	if hiddenp(str1, str2) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune(10)
}
