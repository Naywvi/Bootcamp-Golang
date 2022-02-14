package piscine

func CountIf(f func(string) bool, tab []string) int {
	SLICEBOOL := []bool{}
	count := 0
	separate := ""
	for i := 0; i < len(tab); i++ {
		separate = tab[i]
		SLICEBOOL = append(SLICEBOOL, f(separate))
		if SLICEBOOL[i] == true {
			count++
		}
	}
	return count
}
