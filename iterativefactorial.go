package piscine

func IterativeFactorial(nb int) int {
	c := nb
	for boucle := 0; boucle < nb; boucle++ {
		if boucle == 0 {
			boucle++
		} else {
			c = c * boucle
		}

	}
	return c
}
