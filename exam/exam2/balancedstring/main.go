package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	var counter int
	var counterC int
	var counterD int
	for _, l := range args {
		if l == 'C' {
			counterC++
		}
		if l == 'D' {
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
