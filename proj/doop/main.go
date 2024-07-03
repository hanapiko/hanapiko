package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
	//"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	int1, overflow := Atoi(args[0])
	if overflow {
		return
	}
	operator := args[1]
	int2, _ := Atoi(args[2])
	fmt.Println(int1)

	// if int1 == 9223372036854775807 {
	// 	return
	// }

	if int1 >= 9223372036854775807 || int1 <= -9223372036854775808 {
		return
	}
	if int2 >= 9223372036854775807 || int2 <= -9223372036854775808 {
		return
	}

	operators := []string{"-", "+", "/", "*", "%"}

	op := false

	for _, ope := range operators {
		if ope == operator {
			op = true
		}
	}

	if !op {
		return
	}
	if int1 < '0' && int1 > '9' {
		return
	}
	if int2 < '0' && int2 > '9' {
		return
	}

	result := 0

	if operator == "+" {
		result = int1 + int2
	}
	if operator == "-" {
		result = int1 - int2
	}
	if operator == "*" {
		result = int1 * int2
	}
	if operator == "/" {
		if int2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = int1 / int2
	}
	if operator == "%" {
		if int2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = int1 % int2
	}
	num := Itoa(result)
	os.Stdout.WriteString(num + "\n")
	// z01.PrintRune('\n')
}

func Atoi(s string) (int, bool) {
	multiplier := 1
	val := 0
	for i, ch := range s {
		if i == 0 && ch == '-' {
			multiplier = -1
			continue
		} else if i == 0 && ch == '+' {
			multiplier = 1
			continue
		}
		if ch < '0' && ch > '9' {
			return 0, true
		}
		val = (val * 10) + (int(ch - '0'))
	}
	result := val * multiplier
	if s != Itoa(result) {
		return 0, true
	}
	return result, false
}

func Itoa(n int) string {
	result := ""
	isneg := false
	if n < 0 {
		isneg = true
		n *= -1
	}

	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n /= 10
		// mod := n % 10
		// startrune := '0'
		// for i := 0; i < mod; i++{
		// 	startrune++
		// }
		// result = string(startrune) + result
		// n = n/10
	}
	if isneg {
		return "-" + result
	}
	return result
}
