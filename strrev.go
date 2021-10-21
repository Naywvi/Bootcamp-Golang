package piscine

func StrRev(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
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
