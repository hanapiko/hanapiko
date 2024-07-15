package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
	//"strconv"
)

func main() {
	args := os.Args[1]
	num := Atoi(args)
	// if err != nil  || num >= 4000 || num <= 0 {
	// 	return
	// }
	if num <= 0 || num >= 4000 {
		fmt.Println("Error")
		os.Exit(0)
	}
	cal, res := Roman(num)
	for _, ch := range cal {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
	// fmt.Println(cal)
	// fmt.Println(res)
}

func Roman(n int) (string, string) {
	rm := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	calculations := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	result := ""
	cals := ""
	for i, ch := range values {
		for n >= ch {
			n = n - ch
			result = result + rm[i]

			if len(cals) > 0 {
				cals = cals + "+"
			}
			cals = cals + calculations[i]
		}
	}
	return cals, result
}

func Atoi(s string) int {
	val := 0
	for _, ch := range s{
		if ch < '0' || ch > '9' {
			fmt.Println("Error")
			os.Exit(0)
		}
		val = val * 10 + int(ch - '0')
	}
	return val
}
