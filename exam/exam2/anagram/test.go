package main

import "fmt"

func smallerNB(arr1, arr2 []rune) []rune {
	if len(arr1) > len(arr2) {
		return arr2
	} else {
		return arr1
	}
}

func SortWord(arr []rune) {
	for i1, l1 := range arr {
		for i2, l2 := range arr {
			if l1 > l2 {
				arr[i1], arr[i2] = arr[i2], arr[i1]
			}
		}
	}
}

func letterfilter(str string) string {
	var res string
	for _, l := range str {
		if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
			res += string(l)
		}
	}
	return res
}

func IsAnagram(str1 string, str2 string) bool {
	arr1 := make([]rune, len(letterfilter(str1)))
	arr2 := make([]rune, len(letterfilter(str2)))
	counter1 := 0
	counter2 := 0
	for _, l1 := range letterfilter(str1) {
		arr1[counter1] = l1
		counter1++
	}
	for _, l2 := range letterfilter(str2) {
		arr2[counter2] = l2
		counter2++
	}
	SortWord(arr1)
	SortWord(arr2)
	for i := 0; i < len(smallerNB(arr1, arr2)); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main() {
	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}
