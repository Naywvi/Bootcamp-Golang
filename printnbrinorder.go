package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune(rune(48))
	}
	if n > 0 {
		var SLICE []int
		PLACER := 0
		ONLEBOUGE := 0
		Change := 0
		for n != 0 {
			PLACER = n % 10
			n /= 10
			SLICE = append(SLICE, PLACER)
		}
		for count := range SLICE {
			ONLEBOUGE = count + 1
		}
		for i := 0; i < ONLEBOUGE-1; i++ {
			for j := 0; j < ONLEBOUGE-i-1; j++ {
				if SLICE[j] > SLICE[j+1] {
					Change = SLICE[j]
					SLICE[j] = SLICE[j+1]
					SLICE[j+1] = Change
				}
			}
		}
		for i := 0; i < ONLEBOUGE; i++ {
			z01.PrintRune(rune(SLICE[i] + 48))
		}
	}
}
