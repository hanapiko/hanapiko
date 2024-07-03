package main

import (
	//"fmt"
	"os"

	 "github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	x :=args[len(args)-1]
		// last := ""
	// for i :=  len(args)-1; i > 0; i-- {
	// 	if last != "" && string(args[i]) == " "{
	// 		break
	// 	}
	// 	last = string(args[i]) + last
	// }
	for _, ch := range x{
		z01.PrintRune(ch)
	}

}