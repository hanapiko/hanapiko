package main
import (
	//"fmt"
	"os"
	 "github.com/01-edu/z01"
)
func main() {
	args := os.Args[1]
	//arg := args[len(args)-1]
		last := ""
	for i :=  len(args)-1; i > 0; i-- {
		if  last != "" && args[i] == ' ' {
			break
		}
		last = string(args[i]) + last
	}
	for _, ch := range last {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}