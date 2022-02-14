// 0	"strconv.Atoi"
// 1	"os.*"
// 2	"append"
// 3	"fmt.*"
// 4	"len"
// 5	"--cast"
// 6	"strings.Split"

package main

import "fmt"

func main() {
	input := " "
	brackets := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	var res string
	for _, l := range input {
		if l == '{' || l == '[' || l == '(' {
			res += string(l)
		}
		if (l == '}' || l == ']' || l == ')') && len(res) > 0 {
			if brackets[l] == rune(res[len(res)-1]) {
				res = res[:len(res)-1]
			} else {
				fmt.Println("Error")
				return
			}
		} else if len(res) == 0 && (l == '}' || l == ']' || l == ')') {
			fmt.Println("Error")
			return
		}
	}
	if res == "" {
		fmt.Println("OK")
		return
	} else {
		fmt.Println("Error")
		return
	}
	// else {
	// 	fmt.Println("Error")
	// 	return
	// }
}
