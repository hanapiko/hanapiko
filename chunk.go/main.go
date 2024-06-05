package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size <= 0{
		z01.PrintRune('\n')
	    return
	}
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunk := (slice[i:end])
		chunks =  append(chunks, chunk)
	}
	
	fmt.Println(chunks)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

