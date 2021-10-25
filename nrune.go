package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	y := 0
	if n < 0 {
		return 0
	}
	for i := range x {
		y = i + 1
		if y == n {
			return x[n-1]
		}
	}
	return 0
}
