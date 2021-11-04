package piscine

func Capitalize(s string) string {
	srune := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		if srune[0] >= 'a' && srune[0] <= 'z' {
			srune[i] = srune[i] - 32
		}
		if srune[i] >= 'a' && srune[i] <= 'z' {
			if srune[i-1] >= 'a' && srune[i-1] <= 'z' || srune[i-1] >= 'A' && srune[i-1] <= 'Z' {

			} else {
				srune[i] = srune[i] - 32
			}
		}
	}
	test := string(srune)
	return test
}
