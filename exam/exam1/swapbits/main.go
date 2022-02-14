package main

import "fmt"

func SwapBits(octet byte) byte {
	return ((octet&0x0F)<<4 | (octet&0xF0)>>4)
}

func main() {
	var b byte = 0100
	var a byte
	a = SwapBits(b)
	for _, l := range string(a) {
		fmt.Print(l)
	}
}
