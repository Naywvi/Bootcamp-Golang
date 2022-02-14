package main

import "fmt"

func noSpace(str string) string {
	var newstr string
	for i, l := range str {
		if l != ' ' {
			newstr += string(l)
		} else if i > 0 && l == ' ' && str[i-1] != ' ' {
			newstr += string(l)
		}
	}
	return newstr
}

func rostring(str string) string {
	newstr := str
	counter := 0
	for _, l := range str {
		if l != ' ' {
			newstr += string(l)
			counter++
		} else if l == ' ' {
			counter++
			break
		}
	}
	newstr = newstr[counter:]
	return newstr
}

func main() {
	str := " "
	str += " "
	newstr := noSpace(str)
	newstr2 := rostring(newstr)
	fmt.Println(newstr2)
}
