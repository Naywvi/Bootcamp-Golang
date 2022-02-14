package main

import (
	"fmt"
)

// 0	"fmt.*"
// 1	"len"
// 2	"os.Args"
// 3	"--cast"
// 4	"make"

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

func nbword(str, charset string) int {
	counter := 0
	for i := range str {
		if i <= len(str)-len(charset) && str[i:i+len(charset)] == charset {
			counter++
		}
	}
	return counter
}

func Split(str, charset string) []string {
	str += charset
	arr := make([]string, nbword(str, charset))
	fmt.Println(str)
	counter1 := 0
	start := 0
	for i := 0; i < len(str); i++ {
		if i <= len(str)-len(charset) && str[i:i+len(charset)] == charset {
			arr[counter1] = str[start:i]
			counter1++
			start = i + len(charset)
		}
	}
	return arr
}
