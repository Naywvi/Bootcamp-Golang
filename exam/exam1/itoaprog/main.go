package main

import "fmt"

func Itoa(nb int) string {
	var str string
	var newstr string
	if nb < 0 {
		nb = nb * -1
		newstr = "-"
		for i := nb; i > 0; i /= 10 {
			r := rune(i%10 + 48)
			str += string(r)
		}
		for i := len(str) - 1; i >= 0; i-- {
			newstr += string(str[i])
		}
	} else if nb == 0 {
		newstr = "0"
	} else if nb > 0 {
		for i := nb; i > 0; i /= 10 {
			r := rune(i%10 + 48)
			str += string(r)
		}
		for i := len(str) - 1; i >= 0; i-- {
			newstr += string(str[i])
		}
	}
	return newstr
}
func main() {
	nb := +234
	fmt.Println(Itoa(nb))
}
