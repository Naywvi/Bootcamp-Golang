package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
	//TODO: <= Ici n à 3x* car n = &b = &a
}

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
	//TODO: <= Ici b = (à la valeur de a sous forme d'adresse "informatique" permettant de le récupéré et de le modifier directement ect ect)
}
