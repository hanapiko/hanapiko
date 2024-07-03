package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// args := os.Args[1:]

	// count := 0
	// for range args {
	// 	count++
	// }
	// result := ""

	// for count != 0{
	// 	mod := count % 10
	// 	startrune := '0'
	// 	for i := 0; i < mod; i++{
	// 		startrune++
	// 	}
	// 	result= string(startrune) + result
	// 	count = count / 10
	// }
	// for _, ch := range result {
	// 	z01.PrintRune(ch)
	// }
	// z01.PrintRune('\n')
	result := len(os.Args[1:])
	z01.PrintRune(rune(result + '0'))
	z01.PrintRune('\n')
}