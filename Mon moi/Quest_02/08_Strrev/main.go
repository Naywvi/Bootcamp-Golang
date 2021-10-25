package main

import "fmt"

func StrRev(s string) string {
	sr := []rune(s)
	for a, b := 0, len(sr)-1; a < b; a, b = a+1, b-1 {
		sr[a], sr[b] = sr[b], sr[a]
	}

	return string(sr)
	//TODO:<= Print le string modifié soit (sr)
}

/*
Ne pas oublier que les valeurs sont placé dans un tableau COMME CECI:
 0  1  2  3  4  5  6  7  8  9  10 11
[H][E][L][L][O][ ][W][O][R][L][D][!]
		Alors ici on convertit le string en rune afin de l'appliquer dans une boucle car > int.
	Dans cette boucle nous mettons deux variables ayant deux valeurs. A et B, A = 0 et B = (Nombre de caractère soit 12) '-1' car [0].
	Puis on dit que A+1 et B-1, a chaque tour.
	Enfin A = B et B =A (a chaque tour les nombres s'inverse)
	Et vue que A<B lorsque A devient supèrieur à B STOP !

*/

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
