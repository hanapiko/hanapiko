package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	//word := "abc"
	args := os.Args[1]
	for _, ch := range args{
		if ch >= 'a' && ch <= 'z' {
			z01.PrintRune('z'-(ch-'a'))
		} else if ch >= 'A' && ch <= 'Z' {
			z01.PrintRune('A'-(ch - 'Z'))
		} else {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}