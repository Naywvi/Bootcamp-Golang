package main

import (
	"fmt"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	nb := "-123"
	data, err := strconv.Atoi(nb)
	if err != nil {
		z01.PrintRune('0')
		return
	}
	fmt.Println(data)
}
