// ## printbits

// ### Instructions

// Write a program that takes a number as argument, and prints it in binary value **without a newline at the end**.

// - If the the argument is not a number, the program should print `00000000`.

// ### Usage :

// ```console
// $ go run . 1
// 00000001$
// $ go run . 192
// 11000000$
// $ go run . a
// 00000000$
// $ go run . 1 1
// $ go run .
// $
// ```
package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		Printbits(0)
		z01.PrintRune('\n')
		return
	}

	Printbits(num)
	z01.PrintRune('\n')
}

func Printbits(n int) {
	str := ""

	if n == 0 {
		str = "00000000"
	}

	for n > 0 {
		digit := n%2
		n /= 2
		str += string(rune(digit + '0'))
	}

	for len(str) < 8 {
		str += string('0')
	}

	rr := []rune(str)

	for i := len(rr)-1; i >= 0; i-- {
		z01.PrintRune(rr[i])
	}
}
