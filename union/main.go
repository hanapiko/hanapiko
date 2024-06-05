package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}
	first := args[0]
	second := args[1]
	str := ""
	for _, g := range first {
		if !Contains(g, str) {
			str += string(g)
		}
	}
	for _, g := range second {
		if !Contains(g, str) {
			str += string(g)
		}
	}
	printString(str)
	z01.PrintRune('\n')
}
func printString(s string) {
	for _, b := range s {
		z01.PrintRune(b)
	}
}
func Contains(r rune, s string) bool {
	for _, h := range s {
		if h == r {
			return true
		}
	}
	return false
}
