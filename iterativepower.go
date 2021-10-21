package piscine

func IterativePower(nb int, power int) int {
	x := *&nb
	if nb == 0 || power == 0 {
		println("NUL DONC ZERO", nb, power)
		return 0
	}
	if nb < 0 || power < 0 {
		println("ZERO", nb, power)
		return 0
	}
	for boucle := 0; boucle < power-1; boucle++ {
		if nb > 0 {
			nb *= nb
			println("je passe ici", nb)
		}
	}
	nb = *&nb
	nb = nb / x
	return nb
}
