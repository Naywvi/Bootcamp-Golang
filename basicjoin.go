package piscine

func BasicJoin(elems []string) string {
	wrld := ""
	for i := 0; i < len(elems); i++ {
		wrld = wrld + elems[i]
	}
	return wrld
}
