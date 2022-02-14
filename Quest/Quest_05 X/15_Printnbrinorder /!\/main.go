package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var SLICE [10]int
	for n != 0 {
		SLICE[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for SLICE[1] > 0 {
			z01.PrintRune(rune(0) + '0')
			SLICE[i]--
		}
	}
	return
}

func main() {
	PrintNbrInOrder(3819)

}

/*
	if n > 0 {
		var SLICE []int
		PLACER := 0
		ONLEBOUGE := 0
		Change := 0

		for n != 0 {
			PLACER = n % 10
			n /= 10
			SLICE = append(SLICE, PLACER) // SLICE[1,9,8 ]
		}

		for count := range SLICE { // 1    || 2 || 3
			ONLEBOUGE = count + 1 // 1     || 3 || 3
		}

		for i := 0; i < ONLEBOUGE-1; i++ { //0<0    || 0<2 || 0<2

			for j := 0; j < ONLEBOUGE-i-1; j++ { // 0< -1 || 0<1 ||0<1

				if SLICE[j] > SLICE[j+1] { // slice[1]>slice[2]
					print("je passe ici")
					Change = SLICE[j]     // Change = slice[1] DONC 9
					SLICE[j] = SLICE[j+1] // slice[1]= slice[2]
					SLICE[j+1] = Change   // slice[3]=9
				}

			}
		}
		for i := 0; i < ONLEBOUGE; i++ {
			z01.PrintRune(rune(SLICE[i] + 48))
		}
	}
*/
