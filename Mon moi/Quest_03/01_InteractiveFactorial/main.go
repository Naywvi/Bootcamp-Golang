package main

import "fmt"

func IterativeFactorial(nb int) int {
	c := nb
	if nb == 0 {
		return 1
	}
	if nb < 1 || nb > 99 {
		return 0
	}
	for boucle := 1; boucle < nb; boucle++ {
		c = c * boucle
		print("je passe ici")
	}
	return c
}

/*
TODO:
EXEMPLE : Factoriel est calculer un nombre par celui qui le précède jusqu'a ce qu'il soit infèrieur à son égal.
Soit 4X1 / 4X2 / 4X2X3 | 3 étant infèrieur à 4 on s'arrête là.
Ici c=nb soit 4. Si nb == 0 return 1. Si nb < 1 ou nb > 99 return 0.
Donc nous calculerons nb x le nombre qui le précède donc la boucle qui cimmence à 1.
Pour éviter tout erreur c = nb. Car sinon 'c' changera car nous le multiplions x la boucle.
*/

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
