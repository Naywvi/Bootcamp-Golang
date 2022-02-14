package main

import "fmt"

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)

}

/*TODO: FIBONNACI
La suite de fibonnaci est une suite logique. Cette suite à pour départ 1. Elle se calcule par le chifre qui la précède.
Soit 1+0 = 1 (oui il y a un zéro cependant il est utiliser seulement pour le commencement.)
Donc:
     |1
 1+0=|1
 1+1=|2
 1+2=|3
 2+3=|5
 3+5=|8
 5+8=|13
8+13=|21
Et la formule pour calculer la suite de fibonnaci est Fn=Fn-1+Fn-2
Soit le return à la fin -> ( Fibonacci(index-1) + Fibonacci(index-2) )
*/

func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}
