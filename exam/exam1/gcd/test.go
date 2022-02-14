package main

import "fmt"

func SmallerNumber(nb1, nb2 int) int {
	if nb1 > nb2 {
		return nb2
	} else {
		return nb1
	}
}

func main() {
	nb1 := 108
	nb2 := 9
	for i := SmallerNumber(nb1, nb2); i > 0; i-- {
		if nb1%i == 0 && nb2%i == 0 {
			fmt.Println(i)
			return
		}
	}
}
