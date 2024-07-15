package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	if isPowerOfTwo(n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n & (n - 1)) == 0
}
