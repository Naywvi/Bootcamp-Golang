package piscine

func IterativeFactorial(nb int) int {
	c := nb
	for boucle := 1; boucle < nb; boucle++ {
		if boucle > 24 || nb < 0 {
			return 0
		} else {
			c = c * boucle
		}
	}
	return c
}
