package main

import (
	"fmt"
)

func printB(b byte) {
	bs := []byte{b}
	for _, n := range bs {
		fmt.Printf("%08b ", n)
	}
}

func main() {
	var b byte = 65

	//fmt.Printf("PrintBinary: %v\n\n", printBinary(int(b)))

	swapped := SwapBits(b)
	printB(b)
	fmt.Print(" -> ")
	printB(swapped)
	println()
}

func SwapBits(octet byte) byte {
	return (octet >> 4) | (octet << 4)
	// bin := printBinary(int(octet))

	// bin = bin[4:] + bin[0:4]

	// return byte(Atoi(bin))
}

func printBinary(nb int) string {
	str := ""
	if nb == 0 {
		str = "00000000"
	}
	for nb > 0 {
		digit := nb % 2
		nb /= 2
		str += string(rune(digit + '0'))
	}
	for len(str) < 8 {
		str += string('0')
	}
	ss := []rune(str)

	out := ""
	for i := len(ss) - 1; i >= 0; i-- {
		out += string(ss[i])
	}

	return out
}

func Atoi(s string) int {
	num := 0

	for _, r := range s {
		num = num*2 + int(r-'0')
	}
	return num
}
