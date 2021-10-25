package piscine

func LastRune(s string) rune {
	x := []rune(s)
	y := 0
	for i := range x {
		y = i + 1
	}
	return x[y-1]
}
