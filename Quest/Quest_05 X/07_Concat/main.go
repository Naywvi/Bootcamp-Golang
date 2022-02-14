package main

import (
	"fmt"
)

func Concat(str1 string, str2 string) string {
	concat := str1 + str2
	return concat
}

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))

}
