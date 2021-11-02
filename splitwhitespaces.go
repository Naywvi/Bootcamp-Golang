package piscine

func SplitWhiteSpaces(s string) []string {
	var listemots []string
	mot := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' && i != len(s)-1 {
			mot = mot + string(s[i])
		} else {
			if i == len(s)-1 {
				mot = mot + string(s[i])
			}
			if mot != "" {
				listemots = append(listemots, mot)
			}
			mot = ""
		}
	}
	return listemots
}
