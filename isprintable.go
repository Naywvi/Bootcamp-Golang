package piscine

func IsPrintable(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= 97 && x <= 122 || x == 32 || x >= 63 && x <= 90 || x >= 33 && x <= 39 {
			compteur2++
		}

	}
	if compteur == compteur2 {
		return true
	}
	return false
}
