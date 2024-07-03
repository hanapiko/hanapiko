package main

import (
	"fmt"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	count := 0
	for _, ch := range str2 {
		if ch == rune(str1[count]) {
			count++
			if count == len(str1) {
				for _, ch := range str1 {
					z01.PrintRune(ch)
				}
				count = 0
			}
	}
}
z01.PrintRune('\n')

}