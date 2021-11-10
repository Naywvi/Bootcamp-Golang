package piscine

func SortIntegerTable(table []int) {
	SLICE := []int{}
	Change := 0
	for i := range table {
		q := (table[len(table)-1-i])
		SLICE = append(SLICE, q)
	}
	for i := range table {
		table[i] = SLICE[i]
	}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if SLICE[j] > SLICE[j+1] {
				Change = SLICE[j]
				SLICE[j] = SLICE[j+1]
				SLICE[j+1] = Change
			}
		}
	}
	for i := range table {
		table[i] = SLICE[i]
	}
}
