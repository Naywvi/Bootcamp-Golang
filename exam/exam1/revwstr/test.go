package main

import "fmt"

func main() {
	str := "the time of contempt precedes that of indifference"
	str += " "
	var tempstr string
	var resstr string
	for _, l := range str {
		if l != ' ' {
			tempstr += string(l)
		} else if l == ' ' {
			resstr = tempstr + resstr
			tempstr = " "
		}
	}
	resstr = resstr[1:]
	fmt.Println(resstr)
}
