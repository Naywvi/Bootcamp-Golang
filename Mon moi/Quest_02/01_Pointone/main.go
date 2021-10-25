package main

import "fmt"

func PointOne(n *int) {
	*n = 1
	//TODO: <= *n car nous allons directement la chercher dans la fonction main
}

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
	//TODO: <= Nous envoyons &n car l'adresse de n
}
