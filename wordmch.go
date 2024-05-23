package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	index := 0
	for _, ch := range str2 {
		if ch == rune(str1[index]) {
			index++
		}
		if len(str1) == index {
			fmt.Println(str1)
			return
		}

	}
}
