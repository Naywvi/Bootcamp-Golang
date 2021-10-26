package piscine

func AlphaCount(s string) int {
	k := 0
	for _, x := range s {
		if (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') {
			k++
		}
	}
	return k
}
