package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	SLICEINT := 0
	SLICEINT2 := 1
	count1 := 0
	count2 := 0
	finish := "1"
	for i := len(a) - 2 + 1; i > 0; i-- {
		SLICEINT = a[i]
		SLICEINT2 = a[i-1]
		if f(SLICEINT, SLICEINT2) == 1 {
			count1++
		} else {
			count2++
		}
		if i == 1 {
			for I := range finish {
				SLICEINT = a[I]
				SLICEINT2 = a[I]
				if f(SLICEINT, SLICEINT2) == 1 {
					count1++
				}
			}
		}
	}
	if count1 == len(a)-1 {
		return true
	}
	return false
}
