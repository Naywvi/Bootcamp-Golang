package piscine

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		PrintNbr(a[i])
	}
}
