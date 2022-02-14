package main

import (
	"fmt"
)

func Compare(a, b string) int {
	y := 0
	if a == b {
		return 0
	} else if a > b {
		y++
	} else {
		y = -1
	}
	return y
}

/*TODO:
Ici nous avions directement le même string séparé en deux soit 2xstring = A ET B
Il suffisait juste des les comparées. Si a>b = +1 donc ("Hello!", "Hello!") si on les compares =0
*/

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
