package main

import (
	"fmt"
)

func SortWordArr(array []string) {
	for i1, l1 := range array {
		for i2, l2 := range array {
			if l1 < l2 {
				array[i1], array[i2] = array[i2], array[i1]
			}
		}
	}
}

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
