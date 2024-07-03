package main

import (
	"os"

	"github.com/01-edu/z01"
	// "fmt""
)

// func main() {
// 	fmt.Println(Rot10("my name is hannah"))
// }

// func Rot10(s string) string {
// 	word := []rune{}
// 	result := ""
// 	for _, ch := range s {
// 		if ch >= 'a' && ch <= 'z' {
// 			word = []rune{(ch-'a'+10)%26 +'a'}
// 			result += string(word)
// 		} else if ch >= 'A' && ch <= 'Z' {
// 			word = []rune{(ch-'A'+10)%26 +'A'}
// 			result += string(word)
// 		} else {
// 			result += string(ch)
// 		}
// 	}
// 	return result
// }

func main() {
	args := os.Args[1]
	for _, char := range args {
		if char >= 'a' && char <= 'm' {
			z01.PrintRune(char + 32)
		} else if char >= 'n' && char <= 'z' {
			z01.PrintRune(char - 32)
		}
		if char >= 'A' && char <= 'M' {
			z01.PrintRune(char + 32)
		} else if char >= 'N' && char <= 'Z' {
			z01.PrintRune(char - 32)
		}
	}
	z01.PrintRune('\n')

}
