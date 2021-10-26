package piscine

func Index(s string, toFind string) int {
	for i := range s {
		if i+len(toFind) <= len(s) && string(s[i:i+len(toFind)]) == toFind {
			return i
		}
	}
	return -1
}
