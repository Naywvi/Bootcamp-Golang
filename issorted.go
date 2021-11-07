package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := len(a) - 1; i > 0; i-- {
		if f(a[i], a[i-1]) == -1 || f(a[i-1], a[i]) == -1 {
			return false
		} else if f(a[i], a[i-1]) == 0 {
			return false
		}
	}
	return true
}
