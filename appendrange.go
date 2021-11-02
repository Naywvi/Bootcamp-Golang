package piscine

func AppendRange(min, max int) []int {
	var SLICE []int
	for i := min; i < max; i++ {
		SLICE = append(SLICE, i)
	}
	return SLICE
}
