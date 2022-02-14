package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return nb * IterativePower(nb, power-1)
	}
}

func main() {
	fmt.Println(IterativePower(4, 3))
}

/*
TODO:
	Nous commencons par demander si nb > 24 return 0 sinon chiffre trop grand.
Si nb == 0 alors 0
Si nb < 0 (chiffres neg) alors 0
ELSE : nb x lui même -1 puis retourne au top jusqu'a l'épuisement des puissances.
*/
