package piscine

func IsNumeric(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= '0' && x <= '9' {
			compteur2++
		}

	}
	if compteur == compteur2 {
		return true
	}
	return false
}
