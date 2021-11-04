package piscine

func TrimAtoi(s string) int {
	Value := 0
	degage := 0
	for _, i := range s {
		if i == '-' && Value == 0 {
			degage = 1
		}
		if int(i) >= 48 && int(i) <= 57 {
			Value = (int(i) - 48) + Value*10
		}
	}
	if degage == 1 {
		return -Value
	}
	return Value
}
