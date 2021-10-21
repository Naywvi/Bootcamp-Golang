package piscine

func IterativeFactorial(nb int) int {
	c := nb
	if nb < 0 || nb > 99 {
		return 0
	}
	for boucle := 1; boucle < nb; boucle++ {

		c = c * boucle
	}
	return c
}
