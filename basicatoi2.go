package piscine

func BasicAtoi2(s string) int {
	relou := 0
	zerolens := len(s)
	StringRune := []rune(s)
	for word := 0; word < zerolens; word++ {
		if StringRune[word] < '0' || StringRune[word] > '9' {
			return 0
		} else {
			relou *= 10
			relou += int(StringRune[word]) - '0'
		}
	}
	return relou
}
