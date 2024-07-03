package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	seen := make(map[rune]bool)
	result := ""

	for _, ch := range str1 + str2 {
		if !seen[ch]{
			seen[ch] = true
			result += string(ch)
		}
	}
	for _, ch := range result{
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}