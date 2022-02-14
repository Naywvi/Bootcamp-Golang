package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func main() {
	a1 := []int{985058, 840813, -922648, 676239, 813719, 837710, 275807, -121052}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
