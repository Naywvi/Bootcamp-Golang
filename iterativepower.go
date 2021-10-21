package piscine

func IterativePower(nb int, power int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	for boucle := 0; boucle < 2; boucle++ {
		if nb < 32 {
			nb = nb*nb + 16
		} else {
			nb = nb + 32
		}
	}
	return nb
}
