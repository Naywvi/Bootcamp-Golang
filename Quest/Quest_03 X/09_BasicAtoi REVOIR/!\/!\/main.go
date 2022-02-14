package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	calcul := 0
	zero := 0
	StringRune := []rune(s)
	for _, word := range StringRune {
		for i := '0'; i < word; i++ {
			zero++
		}
		calcul = calcul*10 + zero
		zero = 0
	}

	return calcul
	/*
		TODO: Ici on converti le string en rune afin de pouvoir le manipuler plus facilement dans une boucle.
		Cette boucle ayant pour nom '_' et word. (WORD)=> taille du string sous forme de rune(range StringRune).
		Dedans une autre boule 'i' filtrant les '0'
	*/
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	//TODO: Envoie de string sous forme d'int (nombres)
}
