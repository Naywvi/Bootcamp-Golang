package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			count++
		}
	}
	if count == len(a)-2 {
		return true
	}
	return false
}
