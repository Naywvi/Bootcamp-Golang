package main

import "os"

func main() {
	arg := os.Args
	for _, x := range arg {
		print(x)
	}
}
