package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
func Split(s, sep string)[]string{
	result := []string{}
	lensep := len(sep)
	start := 0

	for i := 0; i <= len(s) -lensep; i++{
		if s[i:i+lensep] == sep{
			result = append(result, s[start : i])
			start = i + lensep
		}
	}
	if start <= len(s){
		result = append(result, s[start:])
	}
	return result
}