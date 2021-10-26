package piscine

func IsAlpha(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' || x == 32 || x >= '0' && x <= '9' {
			compteur2++
		}

	}
	if compteur == compteur2 {
		return true
	}
	return false
}
