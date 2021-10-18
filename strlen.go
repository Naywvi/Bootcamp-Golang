package piscine

func StrLen() {
	var o string = "Hello World!"
	var l int = 0
	for _, s := range o {
		l++
		if s == 1 {
			l++
			s = 12
		}
		if l == 12 {
			println(l)
		}

	}
}
