package main

import (
	"fmt"
)

func IsUpper(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= 'A' && x <= 'Z' {
			compteur2++
		}
	}
	if compteur == compteur2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}
