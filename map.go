package piscine

func Map(f func(int) bool, a []int) []bool {
	SLICEBOOL := []bool{}
	for i := 0; i < len(a); i++ {
		SLICEBOOL = append(SLICEBOOL, f(a[i]))
	}
	return SLICEBOOL
}
