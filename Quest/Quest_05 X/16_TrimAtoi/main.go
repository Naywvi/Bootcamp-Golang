package main

import (
	"fmt"
)

func TrimAtoi(s string) int {
	Value := 0
	degage := 0
	for _, i := range s {
		if i == '-' && Value == 0 {
			degage = 1
		}
		if int(i) >= 48 && int(i) <= 57 {
			Value = (int(i) - 48) + Value*10
		}
	}
	if degage == 1 {
		return -Value
	}
	return Value
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
