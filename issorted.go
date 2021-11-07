package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	SLICEINT := 0
	SLICEINT2 := 0
	count1 := 0
	for i := len(a) - 2 + 1; i > 0; i-- {
		SLICEINT = a[i]
		SLICEINT2 = a[i-1]
		if f(SLICEINT, SLICEINT2) == 1 {
			count1++
		}
	}
	if count1 == len(a)-1 {
		return true
	}
	return false
}
