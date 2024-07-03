package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	//arg := len(args)
	// z01.PrintRune(rune(arg) + '0')
	// z01.PrintRune('\n')
	Num(len(args))
	z01.PrintRune('\n')
}
func Num(n int) {
	if n < 0 {
		return 
	}
	slice := make([]int, 0)
	for n > 0 {
		slice = append(slice, n%10)
		n /= 10
	}
	for i := len(slice)-1; i >= 0; i--{
		z01.PrintRune(rune(slice[i] + '0'))
	}

}