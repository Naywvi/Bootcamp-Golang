package piscine

func Any(f func(string) bool, a []string) bool {
	SLICEBOOL := []bool{}
	separate := ""
	for i := 0; i < len(a); i++ {
		separate = a[i]
		SLICEBOOL = append(SLICEBOOL, f(separate))
		if SLICEBOOL[i] == true {
			return true
		}
	}
	return false
}
