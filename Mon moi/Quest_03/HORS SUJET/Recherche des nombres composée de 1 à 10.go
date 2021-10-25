package main

import (
	"fmt"
)

func NbEntier(nb int) bool {
	onsenfou := true
	for boucle := 9; boucle > 1; boucle-- {
		y := boucle * boucle
		if y == nb {
			print("true")
		}
		if y < nb {
			print("false")
		}
	}
	return onsenfou
	// N'a pas reussi à Return le bool donc mit en print T/F
}

/*TODO:
Ici nous avons crée une boucle allant de 9>0. On sait qu'un nombre/chiffre entier se multiplie par un même nombre/chiffre afin de donner le même resulta(ex: 4=2x2=4)
Donc y A pour but de multiplier la boucle x2 par elle même a chaque tour. Donc 9x9, 8x8 ....
Et nous verifions si le nb que l'on cherche est égal à y.
Si oui >True>Nombre entier
Si non >False
*/

func main() {
	fmt.Println(NbEntier(5))
	fmt.Println(NbEntier(4))
}
