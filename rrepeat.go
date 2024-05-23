package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	//word := "abc"
	args := os.Args[1]
	for _, ch := range args{
		repeat := 1
		if ch >= 'a' && ch <= 'z' {
			repeat = int(ch - 'a'+1)
		} else if ch >= 'A' && ch <= 'Z' {
			repeat = int(ch - 'A'+1)
		}
		for i := 0; i < repeat; i++ {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}