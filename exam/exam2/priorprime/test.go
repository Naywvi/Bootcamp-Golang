package main

import "fmt"

func main() {
	fmt.Println(Priorprime(14))
	//fmt.Println(isPrime(199))
}

func isPrime(nb int) bool {
	if nb == 0 || nb == 1 {
		return false
	}
	for i := 2; i <= nb; i++ {
		if nb%i == 0 && nb != i {
			return false
		}
	}
	return true
}

func Priorprime(x int) int {
	var res int
	for i := 2; i < 14; i++ {
		if isPrime(i) {
			res += i
		}
	}
	return res
}
