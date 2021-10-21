package piscine

func RecursiveFactorial(nb int) int {
	if nb < 24 {
		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		} else {
			return nb * RecursiveFactorial(nb-1)
		}
	}
	return 0
}
