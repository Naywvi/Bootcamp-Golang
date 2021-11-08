package piscine

func FindNextPrime(nb int) int {
	for a := nb; a >= nb/2; a++ {
		if IsPrime(a) {
			return a
		}
	}
	return 1
}
