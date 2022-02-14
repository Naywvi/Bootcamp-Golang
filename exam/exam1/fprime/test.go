// 0	"strconv.Atoi"
// 1	"os.Args"
// 2	"len"
// 3	"--cast"
// 4	"fmt.*"

package main

import "fmt"

func IsPrime(nb int) bool {
	if nb == 0 || nb == 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 && nb != i {
			return false
		}
	}
	return true
}

func fprime(nb int) {
	if IsPrime(nb) {
		fmt.Println(nb)
		return
	}
	for i := 2; i <= nb; i++ {
		if IsPrime(i) && nb%i == 0 {
			fmt.Print(i)
			fmt.Print("*")
			fprime(nb / i)
			return
		}
	}
}

func main() {
	fprime(5)
}
