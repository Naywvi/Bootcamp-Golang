package piscine

func FindNextPrime(nb int) int {
	for a := nb; a >= nb; a++ {
		if IsPrime(a) {
			return a
		}
	}
	return 1
}
