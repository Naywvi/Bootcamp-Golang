package main

import "fmt"

func AppendRange(min, max int) []int {
	var SLICE []int
	for i := min; i < max; i++ {
		SLICE = append(SLICE, i)
	}
	return SLICE
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
