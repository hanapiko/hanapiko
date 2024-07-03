package main

import "github.com/01-edu/z01"

func main() {
	for i:= 'Z'; i >= 'A'; i-- {
		if i % 2 == 0 {
			z01.PrintRune(i+32)
		} else {
			z01.PrintRune(i)
		}
	}
}