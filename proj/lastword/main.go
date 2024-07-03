package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return // Exit if there are not exactly 2 arguments
	}
	args := os.Args[1] // Get the argument passed to the program
	check := false     // Initialize a flag to check for non-space characters

	// Check for any non-space character in the argument
	for _, c := range args {
		if c != ' ' {
			check = true
			break // Exit loop if a non-space character is found
		}
	}

	// Print the last word if at least one non-space character was found
	if check {
		last := ""
		for i := len(args) - 1; i >= 0; i-- {
			if string(args[i]) != " " {
				last = string(args[i]) + last
			} else if last != "" {
				break
			}
		}
		for _, ch := range last {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
