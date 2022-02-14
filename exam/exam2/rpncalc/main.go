package main

import (
	"fmt"
	"strconv"
)

// 0	"strconv.Atoi"
// 1	"os.*"
// 2	"append"
// 3	"fmt.*"
// 4	"len"
// 5	"--cast"
// 6	"strings.Split"

func RPNcalc(arr []string) {
	var res int
	//fmt.Println(arr)
	for i, l := range arr {
		if l == "+" || l == "-" || l == "*" || l == "/" || l == "%" {
			// fmt.Println(arr)
			nb1, err1 := strconv.Atoi(string(arr[i-2]))
			nb2, err2 := strconv.Atoi(string(arr[i-1]))
			if err1 != nil || err2 != nil {
				fmt.Println("Error")
				return
			}
			if l == "+" {
				res = nb1 + nb2
			} else if l == "-" {
				res = nb1 - nb2
			} else if l == "/" {
				res = nb1 / nb2
			} else if l == "*" {
				res = nb1 * nb2
			} else if l == "%" {
				res = nb1 % nb2
			}
			arr[i] = Itoa(res)
			arr = append(arr[:i-2], arr[i:]...)
			if len(arr) == 1 {
				fmt.Println(arr[0])
				return
			}
			RPNcalc(arr)
			return
		}
	}
	fmt.Println("Error")
	return
	//os.Exit(0)
}

func Itoa(nb int) string {
	var res string
	if nb == 0 {
		return "0"
	}
	for i := nb; i > 0; i /= 10 {
		res += string(i%10 + 48)
	}
	return res
}

func main() {
	input := "1 3 * 2 -"
	var arr = []string{}
	for _, l := range input {
		if l != ' ' {
			arr = append(arr, string(l))
		}
	}
	RPNcalc(arr)
}
