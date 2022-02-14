// 	"github.com/01-edu/z01.PrintRune"
// 	"os.*"
// 	"len"
// 	"make"
package main

import "fmt"

func SortArr(arr []int) {
	for i1, l1 := range arr {
		for i2, l2 := range arr {
			if l1 > l2 {
				arr[i1], arr[i2] = arr[i2], arr[i1]
			}
		}
	}
}

func NoZero(arr []int) []int {
	counter := 0
	for _, l := range arr {
		if l == 0 {
			counter++
		}
	}
	arr = arr[0 : len(arr)-counter]
	return arr
}

func uniqOccur(arr []int) bool {
	fmt.Println(arr)
	for i := range arr {
		if i > 0 && arr[i] == arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	str := "lafkefbhbvjyrc;npnfakvvr"
	arr := make([]int, len(str))
	counter1 := 0
	counter2 := 0
	for i := 32; i <= 126; i++ {
		for _, l := range str {
			if l == rune(i) {
				counter2++
			}
		}
		if counter2 != 0 {
			arr[counter1] = counter2
			counter1++
			counter2 = 0
		}
	}
	SortArr(arr)
	if uniqOccur(NoZero(arr)) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
