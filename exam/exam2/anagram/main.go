package main

import "fmt"

func IsAnagram(str1, str2 string) bool {
	var res string
	var tempstr1 string
	var tempstr2 string
	for _, l1 := range str1 {
		if l1 >= 'A' && l1 <= 'Z' || l1 >= 'a' && l1 <= 'z' {
			tempstr1 += string(l1)
		}
	}
	for _, l2 := range str2 {
		if l2 >= 'A' && l2 <= 'Z' || l2 >= 'a' && l2 <= 'z' {
			tempstr2 += string(l2)
		}
	}
	if len(tempstr1) != len(tempstr2) {
		return false
	}
	for _, l1 := range tempstr1 {
		for _, l2 := range tempstr2 {
			if l2 == l1 {
				res += string(l2)
				break
			}
		}
	}
	fmt.Println("tempstr2", tempstr2)
	fmt.Println("tempstr1", tempstr1)
	fmt.Println("res", res)
	if res == tempstr1 {
		return true
	}
	return false
}

func main() {
	test1 := IsAnagram("a  aa", "aab")
	fmt.Println(test1)

	// test2 := IsAnagram("alem", "school")
	// fmt.Println(test2)

	// test3 := IsAnagram("neat", "a net")
	// fmt.Println(test3)

	// test4 := IsAnagram("annamadrigal", "a man and a girl")
	// fmt.Println(test4)
}
