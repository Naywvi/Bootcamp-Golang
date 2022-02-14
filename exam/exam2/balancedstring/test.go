package main

import "fmt"

func main() {
	str := "CCD"
	counter := 0
	counterC := 0
	counterD := 0
	for _, l := range str {
		if l == 'C' {
			counterC++
		} else if l == 'D' {
			counterD++
		}
		if counterC == counterD {
			counter++
			counterC = 0
			counterD = 0
		}
	}

	fmt.Println(counter)
}
