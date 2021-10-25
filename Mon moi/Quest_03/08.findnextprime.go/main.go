package main

import "fmt"

func FindNextPrime(nb int) int {
	if nb%3 == 0 || nb%2 == 0 {
		for i := 0; i < nb; i++ {

			x := 9
			x = nb + 1

			if nb%3 == 0 || nb%2 == 0 {
				return FindNextPrime(x)
			}
			return x
		}
	}
	return nb
}

func main() {
	fmt.Println(FindNextPrime(4))
	fmt.Println(FindNextPrime(5))
}
