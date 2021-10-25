package piscine

func StrRev(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
	/*
		Alors ici on convertit le string en rune afin de l'appliquer dans une boucle car > int.
	Dans cette boucle nous mettons deux variables ayant deux valeurs. A et B, A = 0 et B = (Nombre de caractère soit 12) '-1' car [0].
	Puis on dit que A+1 et B-1
	*/
}

/*
L1 := []rune(s)
L2 := []rune(s)
for i := range s {
L1[i]=L2[len(L1)-1-i]
}
return string(s)
}
*/
