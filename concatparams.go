package piscine

func ConcatParams(args []string) string {
	line := "\n"
	str := ""
	for Boucle := 0; Boucle < len(args); Boucle++ {
		str += (args[Boucle] + line)
	}
	return str
}
