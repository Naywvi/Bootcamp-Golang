package piscine

func BasicAtoi(s string) int {
	relou := 0
	zero := 0
	StringRune := []rune(s)
	for _, word := range StringRune {
		for i := '0'; i < word; i++ {
			zero++
		}
		relou = relou*10 + zero
		zero = 0
	}
	return relou
}
