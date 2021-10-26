package piscine

func Index(s string, toFind string) int {
	r := []rune(s)
	r2 := []rune(toFind)
	for i := range s {
		for j := range toFind {
			if r[i] == r2[j] {
				if r[0] == r2[j] {
					return -1
				}
				return i
			}
		}
	}
	return -1
}
