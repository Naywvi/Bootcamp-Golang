package main

import "fmt"

func StrLen(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
	//TODO: <= Ici 'l' prend 0. Nous créons une boucle de taille 's' et incrémentons 'l++' afin de conter les caractères. Ici nous en avons 12.
}

func main() {
	l := StrLen("Hello World!")
	fmt.Print(l)
	//TODO: <= Ici nous assignons une variable au string et nous décidons d'envoyer directement la valeur du string soit int = string. De ce fait nous devons declarer 'int' après la fonction 'preciser'. Ici c'est StrLen. (nous permettra un return)
}
