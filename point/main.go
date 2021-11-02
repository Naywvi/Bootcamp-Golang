package main

import "fmt"

func setPoint(pointsx *int) {
	*pointsx = 42
}

func main() {
	var pointsx int = 42
	var pointsy int = 21
	setPoint(&pointsx)
	fmt.Printf("x = %d, y = %d\n", pointsx, pointsy)
}
