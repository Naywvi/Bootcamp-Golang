package piscine

func FindNextPrime(nb int) int {
	if nb < 0 || nb == 1 {
		return 2
	}
	for a := nb; a >= nb/2; a++ {
		if IsPrime(a) {
			return a
		}
	}
	return 1
}
