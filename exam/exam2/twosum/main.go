package main

import (
	"fmt"
)

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

func TwoSum(nums []int, target int) []int {
	for i1 := range nums {
		for i2 := range nums {
			if i1 != i2 {
				if nums[i1]+nums[i2] == target {
					res := []int{i1, i2}
					return res
				}
			}
		}
	}
	res := []int{}
	return res
}
