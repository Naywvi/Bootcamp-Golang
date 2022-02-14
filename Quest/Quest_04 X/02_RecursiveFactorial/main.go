package piscine

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb < 24 {
		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		} else {
			return nb * RecursiveFactorial(nb-1)
		}
	}
	return 0
}

/*
TODO: RECURSIVE
Ici nous filtrons comme dans l'exercice précédent.
Ce pendant nous multiplions nb x lui même-1 afin de le calculer par son précédent jusqu'à 0. Et si 0 return 0.
*/

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
