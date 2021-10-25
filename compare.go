package piscine

func Compare(a, b string) int {
	y := 0
	if a == b {
		return 0
	} else if a > b {
		y = +1
	} else {
		y = -1
	}
	return y
}
