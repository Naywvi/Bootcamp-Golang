// "len"
// "--cast"
// "append"
// "github.com/01-edu/z01"

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	PrintMemory(arr)
}

func HexHelper(nb int) rune {
	temp := 'a'
	for i := 10; i < nb; i++ {
		temp++
	}
	return temp
}

func Itoa(nb int) string {
	var res string
	for i := nb; i > 0; i /= 10 {
		res += string(i%10 + 48)
	}
	return res
}

func toHex(nb int) string {
	var res string
	if nb == 0 {
		return "0"
	}
	for i := nb; i > 0; i /= 16 {
		if i%16 < 10 {
			res = string(Itoa(i%16)) + res
		} else if i%16 > 9 {
			res = string(HexHelper(i%16)) + res
		}
	}
	return res
}

func Print(str string) {
	for _, l := range str {
		z01.PrintRune(l)
	}
}

func PrintMemory(arr [10]int) {
	//var res string
	for i := range arr {
		temphex := toHex(arr[i])
		for i := len(temphex); i < 8; i++ {
			temphex += "0"
			if len(temphex) == 4 {
				temphex += " "
			}
		}
		if i > 0 && i != 4 && i != 8 {
			z01.PrintRune(' ')
		}
		if i == 4 || i == 8 {
			z01.PrintRune(10)
		}
		Print(temphex)
	}
	z01.PrintRune(10)
	for _, l := range arr {
		if l >= 32 && l <= 126 {
			z01.PrintRune(rune(l))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune(10)
}
