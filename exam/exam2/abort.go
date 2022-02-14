package main

import (
	"fmt"
)

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
func Abort(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}
	arrout := Sort(arr)
	return arrout[2]
}

func Sort(arr [5]int) [5]int {
	for i1 := range arr {
		for i2 := range arr {
			if arr[i1] < arr[i2] {
				arr[i1], arr[i2] = arr[i2], arr[i1]
			}
		}
	}
	return arr
}
