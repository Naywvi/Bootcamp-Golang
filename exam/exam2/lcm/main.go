package main

import (
	"fmt"
)

func main() {
	fmt.Println(Lcm(8, 6))
}

func BiggerNb(nb1, nb2 int) int {
	if nb1 > nb2 {
		return nb1
	}
	return nb2
}

func Lcm(first, second int) int {
	var res int
	for i := BiggerNb(first, second); i <= first*second; i++ {
		if i%first == 0 && i%second == 0 {
			res = i
			return res
		}
	}
	return res
}
