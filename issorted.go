package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i], a[i+1]) == -1 {
			return false
		} else if f(a[i], a[i+1]) == 0 {
			return false
		}
	}
	return true
}
