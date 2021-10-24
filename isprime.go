package piscine

func IsPrime(nb int) bool {
	onsenfou := true
	for boucle := 9; boucle > 1; boucle-- {
		y := boucle * boucle
		if y == nb {
			return false
		}
		if y < nb {
			return true
		}
	}
	return onsenfou
}
