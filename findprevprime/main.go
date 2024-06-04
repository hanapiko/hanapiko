package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 0; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(nb int) bool {
	for i := 2; i <= nb/2; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
