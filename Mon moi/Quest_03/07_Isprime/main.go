package main

import "fmt"

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb == 1 {
			return false
		}
		if nb == 2 || nb == 3 {
			return true
		}
		if nb%3 == 0 || nb%2 == 0 {
			return false
		}
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

/*TODO:
On sait qu'un nombre premmier et non pas un chiffre n'est pas divisable. Donc il n'est pas factorisable. Alors c'est un nombre composée. On sait aussi qu'un nombre premier ne retourne jamais un float.
*/

func main() {
	fmt.Println(IsPrime(370676))
	fmt.Println(IsPrime(370676))
}
