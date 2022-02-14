package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	k := 0
	for _, x := range s {
		// ( (rune(x)) >=65 && (rune(x)) <=90) correspond a (x >= 'A' && x <= 'Z') <= AUTRE MANIERE DE DEFINIR UN INTERVAL
		if (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
			k++
		}
	}
	return k
}

/*TODO:
Ici on nous demande de compté toute les lettres d'un string.
Donc : for _, x := range s (Soit pour la boucle _, x(<= va parcourir le string) qui est de taille s)
Puis vérifier (x >= 'A' && x <= 'Z') Donc de A à Z et (x >= 'a' && x <= 'z') de a à z.
Et = 10 car "Hello World"
*/
func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
