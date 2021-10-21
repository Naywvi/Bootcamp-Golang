package piscine

func IterativeFactorial(nb int) int {
	c := nb
	for boucle := 1; boucle < nb; boucle++ {
		c = c * boucle
	}
	return c
}
