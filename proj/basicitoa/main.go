// package main

// import "fmt"

//	func main() {
//		num := 7
//		startrune := '0'
//		for i:= 0; i < num; i++{
//			fmt.Printf("%v,%T OF i \n",i,i)
//			startrune++
//			fmt.Printf("%v,%T OF startrune \n",startrune,startrune)
//		}
//		// fmt.Println(string(48))
//	}
package main

//import "fmt"
import "github.com/01-edu/z01"

func main() {
	num := 123
	result := ""
	
	for num != 0 {
		mod := num % 10
		startrune := '0'
		for i := 0; i < mod; i++{
			startrune++
		}
		result = string(startrune) + result
		num = num/10
	}
	for _, ch := range result{
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
	//fmt.Println(result)


	// for i := len(result)-1; i >= 0; i--{
	// 	fmt.Print(string(result[i]))
	// }
}