package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(s int, t string) {
	ans := ""
	ln := 0
	for _, c := range t {
		if c == c {
			ln++
		}
	}
	mx_p := ln
	if s < 0 {
		ans = "-"
		mx_p *= -1
	}
	if ln > 1 {
		for s/mx_p >= ln {
			mx_p *= ln
		}
		for mx_p != 0 {
			ans = ans + string(t[s/mx_p])
			s = s - s/mx_p*mx_p
			mx_p /= ln
		}
		x := map[rune]bool{}
		for _, c := range t {
			if c == '+' || c == '-' {
				ans = "NV"
				break
			}
			if x[c] == false {
				x[c] = true
			} else {
				ans = "NV"
				break
			}
		}
	} else {
		ans = "NV"
	}
	for _, c := range ans {
		z01.PrintRune(c)
	}
}
