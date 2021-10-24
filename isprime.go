package piscine

func IsPrime(nb int) bool {
	onsenfou := true
	for boucle := 2; boucle < 9; boucle++ {
		y := boucle * boucle
		if y == nb {
			x := y == nb
			return x
		} else {
			x := false
			return x
		}
	}
	return onsenfou
}
