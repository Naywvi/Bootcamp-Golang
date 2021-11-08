package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb == 1 || nb == 2 {
			return 2
		}
		if nb == 3 {
			return 3
		}
		for a := nb; a >= nb; a++ {
			if IsPrime(a) {
				return a
			}
		}
		return nb
	} else {
		return 2
	}
}
