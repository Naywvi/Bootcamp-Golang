package piscine

func Sqrt(nb int) int {
	x := 0
	for TableauDePuissance := 10; TableauDePuissance > 0; TableauDePuissance-- {
		x = TableauDePuissance * TableauDePuissance
		if nb == x {
			return TableauDePuissance
		}
	}
	return 0
}
