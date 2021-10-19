package piscine

func BasicAtoi2(s string) int {
	relou := 0
	zero := 0
	StringRune := []rune(s)
	for _, word := range StringRune {
		if word >= 48 && word <= 57 {
			for i := '0'; i < word; i++ {
				zero++
			}
			relou = relou*10 + zero
			zero = 0
		}

	}
	return relou
}
