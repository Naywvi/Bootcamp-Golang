package piscine

func RecursiveFactorial(nb int) int {
	if nb < 24 {
		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		} else {
			return nb * RecursiveFactorial(nb-1)
		}
	}
	print("STOP")
	return 0
}

/*
}*/
// x := 2
// nb = 4
// if nb > x {
// 	x++
// 	nb = nb * x
// 	print(nb)
// 	return nb
// }
// if nb < 1 || nb > 99 {
// 	return 0
// }
// return nb

// if nb == nb {
// 			print(nb)
//}
