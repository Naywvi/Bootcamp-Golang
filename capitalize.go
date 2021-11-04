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
			if srune[i-1] >= 'a' && srune[i-1] <= 'z' || srune[i-1] >= 'A' && srune[i-1] <= 'Z' || srune[i-1] >= '0' && srune[i-1] <= '9' {

			} else {
				srune[i] = srune[i] - 32
			}
		}
		if srune[i] != srune[0] {
			if srune[i] >= 'A' && srune[i] <= 'Z' {
				if srune[i-1] >= 'a' && srune[i-1] <= 'z' || srune[i-1] >= 'A' && srune[i-1] <= 'Z' || srune[i-1] >= '0' && srune[i-1] <= '9' {
					srune[i] = srune[i] + 32
				}
			}
		}
	}
	test := string(srune)
	return test
}
