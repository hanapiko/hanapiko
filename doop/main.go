package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		//fmt.Println("usage : go run . int op int")
		return
	}
	int1 := args[0]
	operator := args[1]
	int2 := args[2]
	result := 0

	if operator == "+" {
		result = Atoi(int1) + Atoi(int2)
	} else if operator == "-" {
		result = Atoi(int1) - Atoi(int2)
	} else if operator == "*" {
		result = Atoi(int1) * Atoi(int2)
	} else {
		return
	}
	fmt.Println(result)
}

func Atoi(s string) int {
	multiplier := 1
	val := 0
	for i, ch := range s {
		if i == 0 && ch == '-'{
			multiplier = -1
			continue
		} else if i == 0 && ch == '+' {
			multiplier = 1
			continue
		}
		if ch < '0' || ch > '9' {
			os.Exit(1)
		}
		val = val*10 + (int(ch - '0'))
	}
	return val * multiplier
}