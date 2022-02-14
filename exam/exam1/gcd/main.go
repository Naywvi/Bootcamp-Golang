package main

import (
	"fmt"
	"os"
	"strconv"
)

func BiggerNb(nb1 int, nb2 int) int {
	if nb1 > nb2 {
		return nb1
	}
	return nb2
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println()
		return
	} else {
		str1 := os.Args[1]
		str2 := os.Args[2]
		flag := 0
		nb1, err1 := strconv.Atoi(str1)
		nb2, err2 := strconv.Atoi(str2)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		for i := 1; i < (BiggerNb(nb1, nb2)); i++ {
			if nb1%i == 0 && nb2%i == 0 {
				flag = i
			}
		}
		fmt.Println(flag)
	}
}
