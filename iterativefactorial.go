package piscine

func IterativeFactorial(nb int) int {
	c := nb
	for boucle := 1; boucle < nb; boucle++ {
		if nb > 24 || nb <= 0 {
			return 0
		} else if boucle > 0 || boucle > 24 {
			c = c * boucle
		}
	}
	return c
}
