package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb == 1 {
			return 2
		}
		if nb == 2 {
			return 2
		}
		if nb == 3 {
			return 3
		}
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				return FindNextPrime(nb + 1)
			}
		}
		return nb
	} else {
		return 2
	}
}
