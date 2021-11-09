package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	count := 0
	for i := len(base) - 1; i >= 0; i-- {
		s := rune(base[i])
		for _, x := range base {
			if s == x {
				count++
				if count > 1 {
					z01.PrintRune('N')
					z01.PrintRune('V')
					return
				}
			}
		}
		if base[i] == '-' || base[i] == '+' {
			return
		}
		count = 0
	}
	if nbr < 0 {
		IndexNEG(CalculeNEG(nbr, base), nbr, base)
	} else {
		Indexx(Calcule(nbr, base), nbr, base)
	}
}

func traitement(SLICE []int) []int {
	SLICE2 := []int{}
	for i := len(SLICE) - 1; i >= 0; i-- {
		SLICE2 = append(SLICE2, SLICE[i])
	}
	return SLICE2
}

func Calcule(nb int, s string) []int {
	SLICE := []int{}
	if nb > 0 {
		nb *= -1
	}
	for nb != 0 {
		nbr := nb
		nb %= len(s)
		nbr /= len(s)
		nb *= -1
		SLICE = append(SLICE, nb)
		nb = nbr
	}
	return traitement(SLICE)
}

func CalculeNEG(nb int, s string) []int {
	SLICE := []int{}
	for nb != 0 {
		nbr := nb
		nb %= len(s)
		nbr /= len(s)
		nb *= -1
		SLICE = append(SLICE, nb)
		nb = nbr
	}
	return traitement(SLICE)
}

func Indexx(SLICE []int, nb int, s string) {
	sSLICE := make([]rune, len(SLICE))
	srune := []rune(s)
	for i := 0; i < len(SLICE); i++ {
		x := SLICE[i]
		sSLICE[i] = srune[x]
		Printer(sSLICE[i])
	}
}

func IndexNEG(SLICE []int, nb int, s string) {
	sSLICE := make([]rune, len(SLICE))
	srune := []rune(s)
	z01.PrintRune('-')
	for i := 0; i < len(SLICE); i++ {
		x := SLICE[i]
		sSLICE[i] = srune[x]
		Printer(sSLICE[i])
	}
}

func Printer(pr rune) {
	z01.PrintRune(pr)
}
