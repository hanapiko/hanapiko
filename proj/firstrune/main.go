package main

import (
	"github.com/01-edu/z01"


)

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	//method 1
	// first := s[0]
	// return rune(first)
	 
	// last := s[len(s)-1]
	// return rune(last)
    
	//method2
	// return rune(s[0])
	// return rune(s[len(s)-1])
    
	//method3
	a := []rune(s)
	return a[0]

	// a := []rune(s)
	// return a[len(s)-1]

}
