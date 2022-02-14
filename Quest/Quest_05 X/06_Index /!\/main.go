package main

import (
	"fmt"
)

func Index(s string, toFind string) int {
	for i := range s {
		if i+len(toFind) <= len(s) && string(s[i:i+len(toFind)]) == toFind {
			return i
		}
	}
	return -1
}

/*TODO:
Je dois verifier la case [?] du second string toFind. S'il n'existe pas dans le string 's' return -1.
______||
i+len(toFind) <= len(s)
Soit i le nombre de fois qu'il a tourné car =range de s + la taille de tofind est infèrieur à la taille de s.
______||
string(s[i:i+len(toFind)]) == toFind
Dans le string S on cible >[] qui va compté de 0 à 1 et s'arreteras à 2.
Entre i et i soit (i:i) + len de toFind donc 6+1 ++ toFind
______||
Sinon return -1
*/
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))

}
