package piscine

func IsPrime(nb int) bool {
	nn := false
	for boucle := 2; boucle < 9; boucle++ {
		y := boucle * boucle
		if y == nb {
			a := true
			return a
		}
		if y < boucle {
			b := false
			return b
		}
	}
	return nn
}
