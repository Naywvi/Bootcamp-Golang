package piscine

import "fmt"

func IsLower(s string) bool {
	compteur2 := 0
	compteur := 0
	for _, x := range s {
		compteur++
		if x >= 'a' && x <= 'z' {
			compteur2++
		}
	}
	if compteur == compteur2 {
		return true
	}
	return false
}

/*TODO:
 */
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
