package piscine

func IterativeFactorial(nb int) int {
	c := nb

	for boucle := 1; boucle < nb; boucle++ {
		if nb < 0 || nb > 99 {
			c = c * boucle
		}
	}
	return c
}
