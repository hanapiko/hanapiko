package main

import (
	// "strconv"

	"strconv"

	"github.com/01-edu/z01"
	// "github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{50000, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	str := strconv.Itoa(result)
	Str(str)
	// Num(result)
	z01.PrintRune('\n')
	// fmt.Println(str)
}

//	func ReduceInt(a []int, f func(int, int) int){
//		result := a[0]
//		for i := 1; i < len(a); i++ {
//			result = f(result, a[i])
//		}
//		Num(result)
//		z01.PrintRune('\n')
//	}
// func Num(n int) {
// 	if n < 0 {
// 		return
// 	}
// 	slice := make([]int, 0)
// 	for n > 0 {
// 		slice = append(slice, n%10)
// 		n /= 10
// 	}
// 	for i := len(slice) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(slice[i] + '0'))
// 	}
// }

func Str(s string) {
	for _, s := range s {
		z01.PrintRune(s)
	}
}
