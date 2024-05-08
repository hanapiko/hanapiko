package myfunctions

import (
	"fmt"
	"strconv"
)
// Hex function takes a slice of strings as input and converts hexadecimal strings to decimal numbers.
func Hex(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			number, err := strconv.ParseInt(s[i-1], 16, 64)// used to converting string to integer
			if err != nil {
				fmt.Println(err)
			}
			s[i-1] = fmt.Sprint(number)
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
// Bin function takes a slice of strings as input and converts hexadecimal strings to decimal numbers.
func Bin(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(bin)" {
			number, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				fmt.Println(err)
			}
			s[i-1] = fmt.Sprint(number) // converting number back to string
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}