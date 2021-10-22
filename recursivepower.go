package piscine

func ReccursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return nb * ReccursivePower(nb, power-1)
	}
}
