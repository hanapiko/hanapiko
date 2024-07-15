package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		// fmt.Println("usage : go run . int op int")
		return
	}
	int1 := args[0]
	operator := args[1]
	int2 := args[2]
	result := 0

	num1 := Atoi(int1)
	num2 := Atoi(int2)

	if num1 >= 9223372036854775805 || num1 <= -9223372036854775805 {
		os.Exit(0)
		return
	}

	if num2 >= 9223372036854775805 || num1 <= -9223372036854775805 {
		os.Exit(0)
		return
	}

	if operator != "*" && operator != "/" && operator != "+" && operator != "-" && operator != "%" {
		os.Exit(0)
		return
	}

	switch operator {
	case "+":
		result = num1 + num2 
	case  "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0")
			 return
		}
		result = num1 / num2
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0")
			 return
		}
		result = num1 % num2
	}
	
    os.Stdout.WriteString(itoa(result)+ "\n")
	//os.Stdout.WriteString("\n")

}

func itoa(n int) string {
	sign := ""
	res := ""

	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return sign + res
}


func Atoi(s string) int {
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
		if ch < '0' || ch > '9' {
			os.Exit(0)
		}
		val = val*10 + (int(ch - '0'))
	}

	return val * multiplier
}
