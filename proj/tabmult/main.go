package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	first := args[0]
	num := 0
	for _, r := range first {
		if r >= '0' && r <= '9' {
			num = num*10 + int(r - '0')
		} else {
			os.Stdout.WriteString("character invalid\n")
			return
		}
	}
	result := 0
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		//Num(num)
		Itoa(num)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		result = i * num
		//Num(result)
		Itoa(result)
		z01.PrintRune('\n')
	}
	// z01.PrintRune('\n')
}

func Itoa(n int) {
	digit := make([]int, 0)
	str := ""
	for n > 0 {
		digit = append(digit, n%10)
		n /= 10
	}
	for i := len(digit)-1; i >= 0; i-- {
		str += string((rune(digit[i] + '0')))
	}
	//return str
	for _, ch := range str {
		z01.PrintRune(ch)
	}
}

// func Num(n int) {
// 	if n < 0 {
// 		return
// 	}
// 	digit := make([]int, 0)
// 	for n > 0 {
// 		digit = append(digit, n%10)
// 		n /= 10
// 	}
// 	for i := len(digit)-1; i >= 0; i-- {
// 		z01.PrintRune(rune(digit[i] + '0'))
// 	}
// }