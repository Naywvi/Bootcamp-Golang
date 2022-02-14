package main

import (
	"fmt"
)

func Itoa(n int) string {
	var res string
	var result string
	if n == 0 {
		return "0"
	} else if n > 0 {
		for i := n; i > 0; i /= 10 {
			counter := rune(i%10 + 48)
			res += string(counter)
		}
		for i := len(res) - 1; i >= 0; i-- {
			result += string(res[i])
		}
	} else if n < 0 {
		n *= -1
		for i := n; i > 0; i /= 10 {
			counter := rune(i%10 + 48)
			res += string(counter)
		}
		res += "-"
		for i := len(res) - 1; i >= 0; i-- {
			result += string(res[i])
		}
	}
	return result
}
func main() {
	nb := -234
	fmt.Println(Itoa(nb))
}
