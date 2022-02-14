package main

import "github.com/01-edu/z01"

func setPoint(pointsx *int, pointsy *int) {
	*pointsx = 42
	*pointsy = 21
	nb1 := *pointsx
	nb2 := *pointsy
	SLICE1 := []int{}
	SLICE2 := []int{}
	for boucle1 := 0; boucle1 < 2; boucle1++ {
		if nb1 > 10 || nb2 > 10 {
			nb1 = nb1 % 10
			SLICE1 = append(SLICE1, nb1)
			nb2 = nb2 % 10
			SLICE2 = append(SLICE2, nb2)
		} else {
			*pointsx /= 10
			SLICE1 = append(SLICE1, *pointsx)
			*pointsy /= 10
			SLICE2 = append(SLICE2, *pointsy)
		}
	}
	z01.PrintRune(rune(120))
	z01.PrintRune(rune(32))
	z01.PrintRune(rune(61))
	z01.PrintRune(rune(32))
	for boucle2 := 1; boucle2 >= 0; boucle2-- {
		slcv1 := rune(SLICE1[boucle2])
		z01.PrintRune(rune(48 + slcv1))
	}
	z01.PrintRune(rune(44))
	z01.PrintRune(rune(32))
	z01.PrintRune(rune(121))
	z01.PrintRune(rune(32))
	z01.PrintRune(rune(61))
	z01.PrintRune(rune(32))
	for boucle3 := 1; boucle3 >= 0; boucle3-- {
		slcv2 := rune(SLICE2[boucle3])
		z01.PrintRune(rune(48 + slcv2))
	}
	z01.PrintRune('\n')
}

func main() {
	pointsx := 0
	pointsy := 0
	setPoint(&pointsx, &pointsy)
}
