package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, x := range s {
		z01.PrintRune(rune(x))
	}
	z01.PrintRune('\n')
	//TODO: Ici nous créons une boucle '_' pour ne pas réutiliser la variable. Et x qui sera égal au string donc il le fera 'pop'. C'est une manière de dire que x = au string sous forme d'int.
	// (:= range s) égale à la taille du string
}

func main() {
	PrintStr("Hello World")
	//TODO: <= Ici nous envoyons un string 's' vers la fonction PrintStr.
}
