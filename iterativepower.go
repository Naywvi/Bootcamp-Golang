package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return nb * IterativePower(nb, power-1)
	}
}

/*
x := *&nb
	if nb == 0 || power == 0 {
		return 0
	}
	if nb < 0 || power < 0 {
		return 0
	}
	if power < nb {
		for boucle := 0; boucle < power-1; boucle++ {
			if nb > 0 {
				nb *= nb
			}
		}
		nb = *&nb
		nb = nb / x
	} else {
		for boucle := 0; boucle < power-1; boucle++ {
			if nb > 0 {
				nb *= nb
			}
		}
		nb = *&nb
		nb = nb / x

	}
	return nb
Rajout de puissance lorsqu'elle est sup au nb. Exemple 3p4 (Rajout 4-3p) CHELOU possibilité que ce soit dans la boucle car il y a -1
*/
