package main

func UltimateDivMod(a *int, b *int) {
	var x int = *a
	*a = *a / *b
	*b = x % *b
}
