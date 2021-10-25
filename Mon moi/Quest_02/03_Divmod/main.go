package main

func Divmod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
	/*TODO: <=Ici *Div pointant vers div de la fonction main aura pour resulta la division de a et b.
	*mod sera le modulo du resulta de $div et b*/
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	Divmod(a, b, &div, &mod)
	println(div)
	println(mod)
	//TODO: <= Ici nous envoyons deux variables soit a et b puis deux variable pointant la variable dans la fonction main via l'*. Permettant ainsi la récupération direct de la valeur et la modifier dans la fonction ciblé.
}
