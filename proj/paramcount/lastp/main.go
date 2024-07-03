package main

import "os"
import "github.com/01-edu/z01"

func main() {
	args := os.Args[1:]
	arg := args[len(args)-1]
	for _, ch := range arg {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
} 