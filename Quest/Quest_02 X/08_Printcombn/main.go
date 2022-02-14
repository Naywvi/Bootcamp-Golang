package main

import "github.com/01-edu/z01"

func PrintComb1(n int) { //--1
	a := 0
	for boucle := 0; boucle < 10; boucle++ {
		a++
		if a < 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(44))
			z01.PrintRune(rune(32))
		} else if a == 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune('\n'))
			break
		}
	}
}
func PrintComb2(n int) { //--2
	a := 0
	b := 0
	for boucle := 0; boucle < 100; boucle++ {
		a++
		if a == 10 {
			b++
			a = 0
		}
		if a == 9 && b == 8 {
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune('\n'))
			break
		}
		if a < 9 && b < 9 {
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(44))
			z01.PrintRune(rune(32))
		}
	}
}

func PrintComb3(n int) { //--3
	a := 1
	b := 1
	c := 0
	for boucle := 0; boucle < 1000; boucle++ {
		a++
		if a == 10 {
			b++
			a = 0
		}
		if b == 10 {
			c++
			b = 0
		}
		if a == 9 && b == 8 && c == 7 {
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune('\n'))
			break
		}
		if a < 9 && b < 9 && c < 9 {
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(44))
			z01.PrintRune(rune(32))
		}
	}
}

func PrintComb4(n int) { //--4
	a := 0
	b := 1
	c := 2
	d := 2
	for i := 0; i < 10000; i++ {
		d++
		if d == 10 {
			c++
			d = 0
		}
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
			a++
			b = 0
		}
		if a == 9 && b == 8 && c == 9 && d == 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune('\n'))
			break
		}

		if a*10+b < c*10+d {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		}
	}
}

func PrintComb5(n int) { //--5
	a := 0
	b := 1
	c := 2
	d := 3
	e := 3
	for i := 0; i < 100000; i++ {
		e++
		if e == 10 {
			d++
			e = 0
		}
		if d == 10 {
			c++
			d = 0
		}
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
			a++
			b = 0
		}

		if a == 5 && b == 6 && c == 7 && d == 8 && e == 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(48 + e))
			z01.PrintRune(rune('\n'))
			break
		}

		if a*10+b < c*10+d && d < e {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(48 + e))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		}
	}
}

func PrintComb6(n int) { //--6
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	for i := 0; i < 1000000; i++ {
		f++
		if f == 10 {
			e++
			f = 0
		}
		if e == 10 {
			d++
			e = 0
		}
		if d == 10 {
			c++
			d = 0
		}
		if c == 10 {
			b++
			c = 0
		}
		if b == 10 {
			a++
			b = 0
		}

		if a == 4 && b == 5 && c == 6 && d == 7 && e == 8 && f == 9 {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(48 + e))
			z01.PrintRune(rune(48 + f))
			z01.PrintRune(rune('\n'))
			break
		}

		if a*10+b < c*10+d && c*10+d < f*10+e {
			z01.PrintRune(rune(48 + a))
			z01.PrintRune(rune(48 + b))
			z01.PrintRune(rune(48 + c))
			z01.PrintRune(rune(48 + d))
			z01.PrintRune(rune(48 + e))
			z01.PrintRune(rune(48 + f))
			z01.PrintRune(rune(44))
			z01.PrintRune(32)
		}
	}
}

/*
func PrintComb7(n int) { //--7
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
}
func PrintComb8(n int) { //--8
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0
}
func PrintComb9(n int) { //--9
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0
	i := 0
}
*/
func PrintCombN(n int) {
	if n == 1 {
		PrintComb1(n)
	} else if n == 2 {
		PrintComb2(n)
	} else if n == 3 {
		PrintComb3(n)
	} else if n == 4 {
		PrintComb4(n)
	} else if n == 5 {
		PrintComb5(n)
	} else if n == 6 {
		PrintComb6(n)
	} /*else if n == 7 {
		PrintComb7(n)
	} else if n == 8 {
		PrintComb8(n)
	} else if n == 9 {
		PrintComb9(n)
	} else {
		return
	}
	*/
}
func main() {
	PrintCombN(6)

}

/*
//len(SLICE)*n = nb boucle

func PrintAdn(adnr int, n int, lensl int) { //print
	println(lensl, adnr, n)
	if n > 1 {
		z01.PrintRune(rune(48 + adnr))
		z01.PrintRune(rune(44))
		z01.PrintRune(rune(32))
	} else {
		z01.PrintRune(rune(48 + adnr))
		z01.PrintRune(rune(44))
		z01.PrintRune(rune(32))
	}

}

func PrintCombN(n int) {
	adn := n
	SLICE := []int{}
	lenn := func(n int) { // ici
		ns := 0
		for n != 0 {
			if n != 0 {
				ns = n % 10
				n /= 10
				SLICE = append(SLICE, ns)
			} else {
				return
			}
		}
	}
	lenn(n)
	lensl := len(SLICE)      // taille de la len
	x := func(adn int) int { //
		adnmax := 10
		if adn < adnmax {
			for adn < adnmax {
				PrintAdn(adn, n, lensl)
				adn++
			}
			return adn
		} else {
			PrintAdn(adn, n, lensl)
		}
		return adn
	}
	(x(adn)) // calcule
}
*/
