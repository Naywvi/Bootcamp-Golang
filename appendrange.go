package piscine

func AppendRange(min, max int) []int {
	SLICE := []int{}
	for i := min; i < max; i++ {
		SLICE = append(SLICE, i)
	}
	return SLICE
}
