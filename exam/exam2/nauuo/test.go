package main

import "fmt"

func Nauuo(plus, minus, rand int) string {
	if plus > minus+rand {
		return "+"
	} else if minus > plus+rand {
		return "-"
	} else if rand == 0 && plus == minus {
		return "0"
	} else {
		return "?"
	}
}

func main() {
	fmt.Println(Nauuo(50, 43, 20))
	fmt.Println(Nauuo(13, 13, 0))
	fmt.Println(Nauuo(10, 9, 0))
	fmt.Println(Nauuo(5, 9, 2))
}
