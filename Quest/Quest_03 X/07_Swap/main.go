package main

func Swap(a *int, b *int) {
	x := *a
	y := *b
	*b = x
	*a = y
	//TODO: <= Ici nous créons deux variables x,y qui sont égales aux variables pointés. Puis on les swap, pour que la func main récupère les valeurs de a et b.
}
func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	println(a)
	println(b)
	//TODO: <= Ici nous envoyons deux variables à func Swap.
}
