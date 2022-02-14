package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	x := []rune(s)
	y := 0
	for range x {
		y++
	}
	return x[y-1]
}

/*TODO:
/!\ Ici nous devons soutirer une "rune" à la fonciton. /!\
De ce fait nous transformons le string en rune. Puis nous assignons une variable 0.
Cette variable nous permettra de compter combien de fois elle tournera dans la boucle qui aura pour taille le string.
Puis nous allons returne la rune plus le nombre que nous avons compter avec y-1.
*/

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
