package main

import "fmt"

func MakeRange(min, max int) []int {
	t := []int{}
	if max < min {
		return t
	}
	a := max - min
	s := make([]int, a)
	for i := 0; i < a; i++ {
		s[i] = min + i
	}
	return s
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
