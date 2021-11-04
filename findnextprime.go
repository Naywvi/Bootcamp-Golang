package piscine

func FindNextPrime(nb int) int {
	if nb < 0 || nb == 1 || nb == 3 {
		return 2
	} else if nb == 3 {
		return 3
	}
	if nb%3 == 0 || nb%2 == 0 {
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				return FindNextPrime(nb + 1)
			}
		}
	}
	return nb
}
