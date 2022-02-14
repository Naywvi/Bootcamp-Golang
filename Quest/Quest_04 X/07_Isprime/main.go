package main

import "fmt"

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb == 1 {
			return false
		} else if nb == 2 || nb == 3 {
			return true
		} else if nb%3 == 0 || nb%2 == 0 {
			return false
		} else {
			nb /= 2
			for a := 2; a < nb; a++ {
				if nb%a == 0 {
					return false
				}
			}
			nb *= 2
		}

		return true
	} else {
		return false
	}
}

/*TODO:
TROUVER UN MOYENS DE DIVISER LE NOMBRE/2 puis faire la boucle par rapport à ça puis le multiplier /2
On sait qu'un nombre premmier et non pas un chiffre n'est pas divisable. Donc il n'est pas factorisable. Alors c'est un nombre composée. On sait aussi qu'un nombre premier ne retourne jamais un float.
*/

func main() {
	fmt.Println(IsPrime(370676))
	fmt.Println(IsPrime(370676))
}
