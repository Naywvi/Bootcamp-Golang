package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	x := []rune(s)
	y := 0
	if n < 0 {
		return 0
	}
	for range x {
		y++
		if y == n {
			return x[n-1]
		}
	}
	return 0
}

/*TODO:
Ici Nous avons fait le même chose que pour l'exercice précédent. Cependant les nombres négatifs retournes 0 soit NULL car = rune soit ASCII.
Donc nous allons chercher la valeur de la dernière lettre du string.
Et nous l'intégrons directement  via x[...]. Sinon 0 > NULL
*/

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
