// package main

// import "fmt"

// func main() {
// 	word := "-2009"
// 	m := 1
// 	result := 0
// 	for i, val := range word {
// 		if i == 0 && val == '-' {
// 			m = -1
// 		} else if i == 0 && val == '+' {
// 			m = 1
// 		} else if val < '0' || val > '9' {
// 			return
// 		} else {
// 			result = (result*10 + int(val-'0'))
// 		}
// 	}
// 	res := result * m
// 	fmt.Println(res)
// }

// package main

// import "fmt"

//	func main() {
//		num := 123
//		result := ""
//		for num != 0 {
//			mod := num%10
//			run := '0'
//			for i := 0; i < mod; i++ {
//				run++
//			}
//			result = string(run) + result
//			num = num/10
//		}
//		fmt.Println(result)
//	}
package main

import "fmt"

func main() {
	word := "123"
	val := 0
	for _, ch := range word {
		val = val * 10 + int(ch - '0')
	}
	fmt.Println(val)
}

