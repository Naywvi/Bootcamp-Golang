package main

import (
	"fmt"
)

func Nauuo(plus, minus, rand int) string {
	if plus == minus && rand == 0 {
		return "0"
	}
	if plus > minus+rand {
		return "+"
	}
	if minus > plus+rand {
		return "-"
	}
	return "?"
}

func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 2))
}
