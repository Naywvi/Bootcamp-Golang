package piscine

func IterativePower(nb int, power int) int {
	if nb == 0 {
		return 0
	}
	if nb < 0 {
		return 0
	}
	for boucle := 0; boucle < 2; boucle++ {
		if nb == 6 {
			nb = nb * 1
			return nb
		}
		if nb < 32 {
			nb = nb*nb + 16
		} else {
			nb = nb + 32
		}
	}
	return nb
}
