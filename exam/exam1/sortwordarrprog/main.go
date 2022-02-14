package main

import "fmt"

func SortWordArr(array []string) {
	for i1 := range array {
		//fmt.Println("after 1 range", array)
		for i2 := range array {
			//fmt.Println("after 2 range", array)
			if array[i1] < array[i2] {
				//fmt.Println("after compare", array)
				array[i1], array[i2] = array[i2], array[i1]
				//fmt.Println("after exchange", array)
			}
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println("final", result)
}
