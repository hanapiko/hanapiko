package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	check := false
	for _, c := range args {
		if c != ' ' {
			check = true
		}
	}
	if check {
		last := ""
		for i := len(args) - 1; i >= 0; i-- {
			if (args[i]) != ' ' {
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
