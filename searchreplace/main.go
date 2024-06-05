package main

import (
	"os"

	"github.com/01-edu/z01"
)

// "fmt"

// func ReplaceAll(str, str1, str2 string) string {
// 	new := ""
// 	for _, ch := range str {
// 		if string(ch) == str1 {
// 			new += str2
// 			continue
// 		}
// 		new += string(ch)
// 	}
// 	return  new
// }

// func main() {
// 	fmt.Println(ReplaceAll("hella there", "a", "o"))
// }

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	first := args[0]
	second := []rune(args[1])
	third := []rune(args[2])
	if len(second) != 1 || len(third) != 1 {
		return
	}
	for _,r := range first {
		if r == second[0] {
			r = third[0]
		}
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}