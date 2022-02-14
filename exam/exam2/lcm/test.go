package main

import "fmt"

func main() {

	fmt.Println(Lcm(6, 9))
}

func isBigger(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func Lcm(first, second int) int {
	var res int
	for i := isBigger(first, second); i <= first*second; i++ {
		if i%first == 0 && i%second == 0 {
			res = i
			return res
		}
	}
	return res
}
