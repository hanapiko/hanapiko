package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(0))
	fmt.Println(FindPrevPrime(6))
}

func FindPrevPrime(nb int) int {
	if !IsPrime(nb){
		nb--
	}
	return nb
}

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb % i == 0 {
			return false
		}
	}
	return true
}
