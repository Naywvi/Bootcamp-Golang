package piscine

func IterativeFactorial(nb int) int {
	c := nb
	save := 0
	for boucle := 1; boucle < nb; boucle++ {
		save++
		c = c * boucle
	}
	return c
}
