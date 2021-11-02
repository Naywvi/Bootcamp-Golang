package piscine

func AppendRange(min, max int) []int {
	SLICE := []int{}
	for i := min; min < max; i++ {
		SLICE = append(SLICE, i)
	}
	return SLICE
}
