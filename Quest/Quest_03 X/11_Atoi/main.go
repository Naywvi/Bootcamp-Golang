package main

import "fmt"

func Atoi(s string) int {
	a := 0
	b := 0
	no := 0
	signe := 0
	count := 0
	for _, r := range s {
		if count == 0 && r == '-' || count == 0 && r == '+' {
			if r == '-' {
				signe = -1
			}
		} else {
			if r < '0' || r > '9' {
				no++
			}
			for i := '0'; i < r; i++ {
				a++
			}
			b = b*10 + a
			a = 0
			if no > 0 {
				b = 0
			}
		}
		count++
	}
	if signe == 1 {
		b = b * 1
	}
	if signe == -1 {
		b = b * -1
	}
	return b
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
