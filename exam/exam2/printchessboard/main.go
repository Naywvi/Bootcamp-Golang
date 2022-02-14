package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// length := os.Args[1]
	// width := os.Args[2]
	nb1 := 4
	nb2 := 5
	var counter int
	black := '#'
	white := ' '
	for i := 1; i <= nb1*nb2; i++ {
		if i%2 != 0 {
			z01.PrintRune(black)
		} else if i%2 == 0 {
			z01.PrintRune(white)
		}
		counter++
		if counter == nb1 && nb1%2 == 0 {
			z01.PrintRune(10)
			temp1 := black
			temp2 := white
			black = temp2
			white = temp1
			counter = 0
		} else if counter == nb1 {
			z01.PrintRune(10)
			counter = 0
		}
	}
}
