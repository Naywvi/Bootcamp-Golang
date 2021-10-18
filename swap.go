package piscine

func Swap(a *int, b *int) {
	var x = *a
	var y = *b
	*a = y
	*b = x
}
