package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	x := []rune(s)
	return x[0]
}

/* TODO:
Ici nous devons rendre une "Rune", et nous recevons un string.
De ce fait nous transformons le string en rune. Il nous est ensuite demandé de print la première
lettre de chaque string envoyé. De ce fait nous choisisons la case [0]
*/
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
