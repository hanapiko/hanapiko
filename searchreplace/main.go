package main

import (
	"fmt"
)

func ReplaceAll(str, str1, str2 string) string {
	new := ""
	for _, ch := range str {
		if string(ch) == str1 {
			new += str2
			continue
		}
		new += string(ch)
	}
	return  new
}

func main() {
	fmt.Println(ReplaceAll("hella there", "a", "o"))
}

