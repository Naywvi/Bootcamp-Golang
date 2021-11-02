package piscine

func MakeRange(min, max int) []int {
	t := []int{}
	if max < min {
		return t
	}
	a := max - min
	s := make([]int, a)
	for i := 0; i < a; i++ {
		s[i] = min + i
	}
	return s
}
