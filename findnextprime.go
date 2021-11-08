package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
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
