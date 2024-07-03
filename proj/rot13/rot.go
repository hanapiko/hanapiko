package main

import "os"

//import "fmt"

// "fmt""

 func main() {


//func Rot10(s string) string {
	//word := []rune{}
	args := os.Args[1:]
	arg := args[0]
	result := ""
	for _, ch := range arg {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch-'a'+13)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch-'A'+13)%26 + 'Z')
		} else {
			result += string(ch)
		}
	}
	os.Stdout.WriteString(result + "\n")
}
	


// func main() {
// 	args := os.Args[1]
// 	for _, char := range args {
// 		if char >= 'a' && char <= 'm' {
// 			z01.PrintRune(char + 32)
// 		} else if char >= 'n' && char <= 'z' {
// 			z01.PrintRune(char - 32)
// 		}
// 		if char >= 'A' && char <= 'M' {
// 			z01.PrintRune(char + 32)
// 		} else if char >= 'N' && char <= 'Z' {
// 			z01.PrintRune(char - 32)
// 		}
// 	}
// 	z01.PrintRune('\n')

// }
