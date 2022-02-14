package main

import "fmt"

func ultimateDivMod(a *int, b *int) {
	x := *a
	*a = *a / *b
	*b = x % *b
	/*TODO: <= Ici la particularité c'est que x tend vers 'a' de la fonction div mod et non pas de la fonction main. Car main ne l'a pas envoyé via '&'.
	Ensuite nous divisons les deux variables de la fonction main et nous la mettons en stock dans *a qui retroune dans main.
	Puis nous faisons la même chose avec *b vers main. Nous voulons le modulo du 'b' de main et le 'a' de divmodm. D'où le stock que nous avons fait au départ.*/
}

func main() {
	a := 13
	b := 2
	ultimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	//TODO: <= Ici nous envoyons directement les valeurs a et b sous forme d'adresse afin de pouvoir les appeler '&' devant la variable nous permet de faire cela.
}
