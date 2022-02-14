package main

import "fmt"

func IsAlpha(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' || x == 32 || x >= '0' && x <= '9' {
			compteur2++
		}

	}
	if compteur == compteur2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
