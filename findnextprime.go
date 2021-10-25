package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb == 2 || nb == 3 {
			return nb
		}
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				return 0
			}
			if nb%3 == 0 || nb%2 == 0 || nb%5 == 0 {
				for i := 0; i < nb; i++ {

					x := 9
					x = nb + 1

					if nb%3 == 0 || nb%2 == 0 || nb%5 == 0 {
						return FindNextPrime(x)
					}
					return x
				}
			}
			return nb
		}
	} else {
		return 0
	}
	return nb
}
