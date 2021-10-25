package main

import (
	"fmt"
)

func Sqrt(nb int) int {
	x := 0
	for TableauDePuissance := 10; TableauDePuissance > 0; TableauDePuissance-- {
		x = TableauDePuissance * TableauDePuissance
		if nb == x {
			return TableauDePuissance
		}
	}
	return 0
}

/*TODO:
Ici on nous demande de trouvé le carrée du nombre donné, soit 4 et 3. En sachant que c'est des chiffres entiers.
Afin de trouver le carrée facilement on multiplie les chiffres de 1 à 9 par leur puissance soit:

A   | 1  | 4  | 9  | 16 | 25 | 36 | 49 | 64 | 81 | <= Le chiffre/nombre calculer à la puissance de celui du dessous soit BXB
    | ------------------------------------------ |
B   | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8  | 9  | <= Le carrée du chiffre/nombre du dessus.

J'ai donc fais une boucle allant de 1 à 9 se multipliant par elle même donc x= boucle*boucle.
Si nb = x alors print. Cependant ici on nous envoie les chiffres 1 par 1. Et si 4 est traité 3 ne le sera pas.
Alors la boucle ne va pas de 1 à 9, mais bien de 9 à 1.
Et si un nombre ne rentre pas dans les conditions alors return 0.
*/

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
