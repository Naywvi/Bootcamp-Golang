package piscine

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb == 1 {
			return false
		}
		if nb == 2 || nb == 3 {
			return true
		}
		if nb%3 == 0 || nb%2 == 0 {
			return false
		}
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
