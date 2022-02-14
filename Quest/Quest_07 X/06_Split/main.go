package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	SLICE := []string{}
	for i := range s {

	}
	return SLICE
}
func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
	/*
		$ go run .
		[]string{"Hello", "how", "are", "you?"}
		$
	*/
}

/*
	indexs := 1
	for i := range s {
		if s[i] == sep[0] && s[i+1] == sep[1] {
			indexs++
		}
	}
	SLICE := make([]string, indexs)
	SLICE[0] = "ok"
	count := 1
	index := 0
	for i := 0; i < len(s)-1; i++ {
		str := string(s[i])
		if count == 1 {
			SLICE[index] = str
			count++
		} else {
			if s[i] == sep[0] && s[i+1] == sep[1] {
				index++
			} else if s[i+1] == sep[0] && s[i+2] == sep[1] {
				str = SLICE[index] + str
				SLICE[index] = str
			} else if s[i-1] == sep[1] && s[i-2] == s[0] && s[i-1] == sep[0] && s[i-2] == s[1] {
				str = SLICE[index] + str
				SLICE[index] = str
			} else if s[i] == sep[1] && s[i-1] == sep[0] {
				//ca degage
			} else {
				str = SLICE[index] + str
				SLICE[index] = str

			}
		}
		if i == len(s)-2 {
			end := string(s[i+1])
			SLICE[index] = str + end
		}
	}
	return SLICE
*/
