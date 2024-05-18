package main

import "fmt"

func main() {
	num := 123
	str := ""
	for num != 0 {
		firstrune := '0'
		mod := num % 10
		for i := 0; i < mod; i++ {
			firstrune++
		}
		str += string(firstrune)
		num = num / 10
	}
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}