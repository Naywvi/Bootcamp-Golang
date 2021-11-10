package piscine

func SortIntegerTable(table []int) {
	SLICE := []int{}
	for i := range table {
		q := (table[len(table)-1-i])
		SLICE = append(SLICE, q)
	}
	for i := range table {
		table[i] = SLICE[i]
	}
}
